// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract Count {

    uint256 public count;

    function increment() public {
        count += 1;
    }
}
