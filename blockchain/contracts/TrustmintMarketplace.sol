// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/interfaces/IERC2981.sol";

/// @title TrustmintMarketplace - Marketplace for trading AI model NFTs
/// @notice Enables listing, buying, and selling of TrustmintNFT tokens
contract TrustmintMarketplace is ReentrancyGuard, Ownable {
    
    struct Listing {
        uint256 listingId;
        uint256 tokenId;
        address nftContract;
        address seller;
        uint256 price;
        bool active;
    }

    // Listing ID counter
    uint256 private _listingIds;

    // Platform fee (in basis points, e.g., 250 = 2.5%)
    uint256 public platformFee = 250; // 2.5%

    // Mapping from listing ID to Listing
    mapping(uint256 => Listing) public listings;

    // Mapping from token ID to listing ID (for quick lookup)
    mapping(address => mapping(uint256 => uint256)) public tokenToListing;

    // Events
    event ListingCreated(
        uint256 indexed listingId,
        uint256 indexed tokenId,
        address indexed seller,
        uint256 price
    );

    event ListingCancelled(uint256 indexed listingId);

    event ModelPurchased(
        uint256 indexed listingId,
        uint256 indexed tokenId,
        address seller,
        address buyer,
        uint256 price
    );

    event PlatformFeeUpdated(uint256 newFee);

    constructor() Ownable(msg.sender) {}

    /// @notice List an NFT for sale
    /// @param nftContract Address of the NFT contract
    /// @param tokenId Token ID to list
    /// @param price Sale price in wei
    /// @return listingId The ID of the created listing
    function listModel(
        address nftContract,
        uint256 tokenId,
        uint256 price
    ) external nonReentrant returns (uint256) {
        require(price > 0, "Price must be greater than zero");
        require(
            IERC721(nftContract).ownerOf(tokenId) == msg.sender,
            "Not the owner of this NFT"
        );
        require(
            IERC721(nftContract).isApprovedForAll(msg.sender, address(this)) ||
            IERC721(nftContract).getApproved(tokenId) == address(this),
            "Marketplace not approved"
        );
        require(
            tokenToListing[nftContract][tokenId] == 0,
            "Token already listed"
        );

        _listingIds++;
        uint256 listingId = _listingIds;

        listings[listingId] = Listing({
            listingId: listingId,
            tokenId: tokenId,
            nftContract: nftContract,
            seller: msg.sender,
            price: price,
            active: true
        });

        tokenToListing[nftContract][tokenId] = listingId;

        emit ListingCreated(listingId, tokenId, msg.sender, price);

        return listingId;
    }

    /// @notice Cancel an active listing
    /// @param listingId The listing ID to cancel
    function cancelListing(uint256 listingId) external nonReentrant {
        Listing storage listing = listings[listingId];
        require(listing.active, "Listing not active");
        require(listing.seller == msg.sender, "Not the seller");

        listing.active = false;
        tokenToListing[listing.nftContract][listing.tokenId] = 0;

        emit ListingCancelled(listingId);
    }

    /// @notice Purchase a listed NFT
    /// @param listingId The listing ID to purchase
    function purchaseModel(uint256 listingId) external payable nonReentrant {
        Listing storage listing = listings[listingId];
        require(listing.active, "Listing not active");
        require(msg.value >= listing.price, "Insufficient payment");
        require(msg.sender != listing.seller, "Seller cannot buy own NFT");

        listing.active = false;
        tokenToListing[listing.nftContract][listing.tokenId] = 0;

        // Calculate fees
        uint256 platformFeeAmount = (listing.price * platformFee) / 10000;
        uint256 royaltyAmount = 0;
        address royaltyReceiver = address(0);

        // Check for ERC2981 royalty support
        if (IERC721(listing.nftContract).supportsInterface(type(IERC2981).interfaceId)) {
            (royaltyReceiver, royaltyAmount) = IERC2981(listing.nftContract).royaltyInfo(
                listing.tokenId,
                listing.price
            );
        }

        // Calculate seller proceeds
        uint256 sellerProceeds = listing.price - platformFeeAmount - royaltyAmount;

        // Transfer NFT to buyer
        IERC721(listing.nftContract).safeTransferFrom(
            listing.seller,
            msg.sender,
            listing.tokenId
        );

        // Transfer funds
        (bool successSeller, ) = payable(listing.seller).call{value: sellerProceeds}("");
        require(successSeller, "Transfer to seller failed");

        if (royaltyAmount > 0 && royaltyReceiver != address(0)) {
            (bool successRoyalty, ) = payable(royaltyReceiver).call{value: royaltyAmount}("");
            require(successRoyalty, "Royalty transfer failed");
        }

        // Refund excess payment
        if (msg.value > listing.price) {
            (bool successRefund, ) = payable(msg.sender).call{value: msg.value - listing.price}("");
            require(successRefund, "Refund failed");
        }

        emit ModelPurchased(
            listingId,
            listing.tokenId,
            listing.seller,
            msg.sender,
            listing.price
        );
    }

    /// @notice Update platform fee (only owner)
    /// @param newFee New platform fee in basis points
    function setPlatformFee(uint256 newFee) external onlyOwner {
        require(newFee <= 1000, "Fee cannot exceed 10%");
        platformFee = newFee;
        emit PlatformFeeUpdated(newFee);
    }

    /// @notice Withdraw accumulated platform fees
    function withdrawFees() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No fees to withdraw");
        (bool success, ) = payable(owner()).call{value: balance}("");
        require(success, "Withdrawal failed");
    }

    /// @notice Get listing details
    /// @param listingId The listing ID
    /// @return Listing struct
    function getListing(uint256 listingId) external view returns (Listing memory) {
        return listings[listingId];
    }

    /// @notice Get total number of listings created
    /// @return uint256 Total listings
    function totalListings() external view returns (uint256) {
        return _listingIds;
    }

    /// @notice Check if token is listed
    /// @param nftContract NFT contract address
    /// @param tokenId Token ID
    /// @return bool True if listed
    function isListed(address nftContract, uint256 tokenId) external view returns (bool) {
        uint256 listingId = tokenToListing[nftContract][tokenId];
        return listingId != 0 && listings[listingId].active;
    }

    // Required to receive NFTs
    function onERC721Received(
        address,
        address,
        uint256,
        bytes memory
    ) public pure returns (bytes4) {
        return this.onERC721Received.selector;
    }
}
