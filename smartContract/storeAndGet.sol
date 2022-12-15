// SPDX-License-Identifier: GPL-3.0

pragma solidity >0.4;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 */
contract Storage {

	mapping (address => mapping(string => string)) examer;
	mapping (string => address) dataIdInStore;
	mapping (string => string) dataVersion;

	function store(string memory _key, string  memory _value, string memory _version) public payable{
		examer[msg.sender][_key] = _value;
		dataIdInStore[_key] = msg.sender;
		dataVersion[_key] = _version;
	}

	function batchStore(string[] memory _keys, string[] memory _values, string[] memory _version) public payable{
		for (uint i = 0; i<_keys.length; i++) {
			examer[msg.sender][_keys[i]] = _values[i];
			dataIdInStore[_keys[i]] = msg.sender;
            dataVersion[_keys[i]] = _version[i];
		}
	}

	function retrieve(string memory _key) public view returns (string memory, string memory, string memory){
		string memory _value = examer[dataIdInStore[_key]][_key];
		string memory _version = dataVersion[_key];
		return (_key, _value, _version);
	}

	function batchRetrieve(string[] memory _keys) public view returns (string[] memory, string[] memory){
		string[] memory values = new string[](_keys.length);
		string[] memory versions = new string[](_keys.length);
		for (uint i=0; i<_keys.length;i++){
			string memory _value = examer[dataIdInStore[_keys[i]]][_keys[i]];
            string memory _version = dataVersion[_keys[i]];
			values[i] = _value;
            versions[i] = _version;
		}
		return (values, versions);
	}
}