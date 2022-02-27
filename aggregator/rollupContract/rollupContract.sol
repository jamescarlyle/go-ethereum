// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;
contract RollupContract {

    struct Rollup {
        bytes32 rollupHash;
        bytes32 parentRollupHash;
        bytes[] transactions;
    }

    Rollup[] public rollups;
 
    /**
     * @dev Store rollup
     * @param rollupHash value to store
     */
    function postRollup(bytes32 rollupHash, bytes32 parentRollupHash, uint256 rollupNumber, bytes[] calldata transactions) external {
        require(rollupNumber == rollups.length, "wrong rollup number");
        Rollup memory rollup = Rollup({
            rollupHash: rollupHash,
            parentRollupHash: parentRollupHash,
            transactions: transactions
        });
        rollups.push(rollup);
    }

    /**
     * @dev Return rollup transactions 
     * @param rollupNumber value to store
     * @return transactions of 'rollup'
     */
    function getTransactions(uint256 rollupNumber) public view returns (bytes[] memory) {
        return rollups[rollupNumber].transactions;
    }
}