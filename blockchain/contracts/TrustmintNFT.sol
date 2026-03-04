// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/common/ERC2981.sol";

/// @title TrustmintNFT - NFT contract for AI models with proof-of-training
/// @notice Each NFT represents a unique AI model with cryptographic proofs
contract TrustmintNFT is ERC721, ERC721URIStorage, ERC2981, Ownable {
    uint256 private _tokenIds;

    // Struct to store model verification data
    struct ModelMetadata {
        string modelHash;      // SHA-256 of the model file
        string datasetHash;    // SHA-256 of the dataset
        bytes32 merkleRoot;    // Merkle root of training proof
        string ipfsCid;        // IPFS CID for model + dataset files
        address creator;       // Original creator/trainer
        uint256 timestamp;     // Block timestamp when minted
        bool verified;         // Verification status
    }

    // Mapping from token ID to model metadata
    mapping(uint256 => ModelMetadata) public modelMetadata;

    // Mapping to track if a model hash has been registered (prevent duplicates)
    mapping(string => bool) public registeredModels;

    // Marketplace contract address (authorized to transfer)
    address public marketplace;

    // Events
    event ModelMinted(
        uint256 indexed tokenId,
        address indexed creator,
        string modelHash,
        string ipfsCid
    );

    event ModelVerified(uint256 indexed tokenId, bool verified);
    event MarketplaceUpdated(address indexed marketplace);

    constructor() ERC721("TrustmintAI", "TMAI") Ownable(msg.sender) {
        // Set default royalty to 5% for contract owner
        _setDefaultRoyalty(msg.sender, 500); // 500 = 5%
    }

    /// @notice Mint a new AI model NFT
    /// @param modelHash SHA-256 hash of the model file
    /// @param datasetHash SHA-256 hash of the dataset
    /// @param merkleRoot Merkle root of the training proof
    /// @param ipfsCid IPFS CID where model and dataset are stored
    /// @param metadataUri IPFS URI for the metadata JSON
    /// @return tokenId The ID of the newly minted NFT
    function mintModel(
        string calldata modelHash,
        string calldata datasetHash,
        bytes32 merkleRoot,
        string calldata ipfsCid,
        string calldata metadataUri
    ) external returns (uint256) {
        require(bytes(modelHash).length > 0, "Model hash cannot be empty");
        require(bytes(ipfsCid).length > 0, "IPFS CID cannot be empty");
        require(!registeredModels[modelHash], "Model already registered");

        _tokenIds++;
        uint256 newTokenId = _tokenIds;

        // Mint NFT to the caller (model creator)
        _safeMint(msg.sender, newTokenId);
        _setTokenURI(newTokenId, metadataUri);

        // Store metadata
        modelMetadata[newTokenId] = ModelMetadata({
            modelHash: modelHash,
            datasetHash: datasetHash,
            merkleRoot: merkleRoot,
            ipfsCid: ipfsCid,
            creator: msg.sender,
            timestamp: block.timestamp,
            verified: true // Auto-verified for now; can add manual verification later
        });

        // Mark model as registered
        registeredModels[modelHash] = true;

        // Set royalty for this token (5% to creator)
        _setTokenRoyalty(newTokenId, msg.sender, 500);

        emit ModelMinted(newTokenId, msg.sender, modelHash, ipfsCid);

        return newTokenId;
    }

    /// @notice Get detailed metadata for a model NFT
    /// @param tokenId The NFT token ID
    function getModelMetadata(uint256 tokenId)
        external
        view
        returns (
            string memory modelHash,
            string memory datasetHash,
            bytes32 merkleRoot,
            string memory ipfsCid,
            address creator,
            uint256 timestamp,
            bool verified
        )
    {
        require(_ownerOf(tokenId) != address(0), "Token does not exist");
        ModelMetadata memory metadata = modelMetadata[tokenId];
        return (
            metadata.modelHash,
            metadata.datasetHash,
            metadata.merkleRoot,
            metadata.ipfsCid,
            metadata.creator,
            metadata.timestamp,
            metadata.verified
        );
    }

    /// @notice Verify a model's authenticity by comparing hashes
    /// @param tokenId The NFT token ID
    /// @param modelHash The hash to verify against
    /// @return bool True if hashes match
    function verifyModel(uint256 tokenId, string calldata modelHash)
        external
        view
        returns (bool)
    {
        require(_ownerOf(tokenId) != address(0), "Token does not exist");
        return keccak256(bytes(modelMetadata[tokenId].modelHash)) == keccak256(bytes(modelHash));
    }

    /// @notice Set the marketplace contract address
    /// @param _marketplace Address of the marketplace contract
    function setMarketplace(address _marketplace) external onlyOwner {
        require(_marketplace != address(0), "Invalid marketplace address");
        marketplace = _marketplace;
        emit MarketplaceUpdated(_marketplace);
    }

    /// @notice Update verification status (only owner)
    /// @param tokenId The NFT token ID
    /// @param verified New verification status
    function setVerificationStatus(uint256 tokenId, bool verified) external onlyOwner {
        require(_ownerOf(tokenId) != address(0), "Token does not exist");
        modelMetadata[tokenId].verified = verified;
        emit ModelVerified(tokenId, verified);
    }

    /// @notice Get total number of minted models
    /// @return uint256 Total supply
    function totalSupply() external view returns (uint256) {
        return _tokenIds;
    }

    // Override required functions
    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721URIStorage, ERC2981)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    function _update(address to, uint256 tokenId, address auth)
        internal
        override(ERC721)
        returns (address)
    {
        return super._update(to, tokenId, auth);
    }
}
