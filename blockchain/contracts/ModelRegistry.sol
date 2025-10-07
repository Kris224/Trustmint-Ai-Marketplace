// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/// @title ModelRegistry - stores provenance information for trained ML models
contract ModelRegistry {
    struct Model {
        uint256 id;
        address owner;
        string modelHash;    // SHA-256 hex (as string)
        string datasetHash;  // SHA-256 hex (as string)
        string ipfsCid;      // Pinata CID / IPFS CID
        bytes32 merkleRoot;  // bytes32 merkle root (0x...)
        uint256 timestamp;
    }

    uint256 public nextId = 1;
    mapping(uint256 => Model) public models;

    event ModelRegistered(uint256 indexed id, address indexed owner, string ipfsCid, bytes32 merkleRoot);

    function registerModel(
        string calldata modelHash,
        string calldata datasetHash,
        string calldata ipfsCid,
        bytes32 merkleRoot
    ) external returns (uint256) {
        uint256 id = nextId++;
        models[id] = Model({
            id: id,
            owner: msg.sender,
            modelHash: modelHash,
            datasetHash: datasetHash,
            ipfsCid: ipfsCid,
            merkleRoot: merkleRoot,
            timestamp: block.timestamp
        });

        emit ModelRegistered(id, msg.sender, ipfsCid, merkleRoot);
        return id;
    }

    function getModel(uint256 id) external view returns (
        uint256, address, string memory, string memory, string memory, bytes32, uint256
    ) {
        Model storage m = models[id];
        return (m.id, m.owner, m.modelHash, m.datasetHash, m.ipfsCid, m.merkleRoot, m.timestamp);
    }
}
