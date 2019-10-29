pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
contract Crud {
  struct User {
    string codmedico;
    string codpaciente;
    string registro;
    string imagen;
  }
  User[] public users;
User[] public aux;
  function create(string memory codmedico, string memory codpaciente, string memory registro, string memory imagen) public {
    users.push(User(codmedico,codpaciente,registro,imagen));
  }

  function read(string memory codpaciente) public returns(User[] memory) {
      delete aux;
          for(uint i = 0; i < users.length; i++) {
      if(compareStrings(users[i].codpaciente,codpaciente)) {
        aux.push(users[i]);
      }
    }
    return(aux);
  }

    function compareStrings (string memory a, string memory b) public view 
       returns (bool) {
  return (keccak256(abi.encodePacked((a))) == keccak256(abi.encodePacked((b))) );

       }
}