pragma solidity >=0.5.0;

contract asset {
    uint64 public TotalSupply;
    string public Name;
    string public Symbol;
    string public Ownership;
    string public Address;
    string public OwnerSocialMedia;
    string public ProofOfTitle;

    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;
    
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approve(address indexed _owner, address indexed _spender, uint256 _value);

    constructor (uint64 _initalSupply, string memory _initialName, string memory _initialSymbol, string memory _initialOwnership, string memory _initialAddress, string memory _initialOwnerSocialMedia, string memory _initialProofOfTitle) public {
        balanceOf[msg.sender] = _initalSupply;

        TotalSupply = _initalSupply;
        Name = _initialName;
        Symbol = _initialSymbol;
        Ownership = _initialOwnership;
        Address = _initialAddress;
        OwnerSocialMedia = _initialOwnerSocialMedia;
        ProofOfTitle = _initialProofOfTitle;
    }

    function transfer(address _to, uint256 _value) public returns (bool success) {
        require(balanceOf[msg.sender] >= _value);
        balanceOf[msg.sender] -= _value;
        balanceOf[_to] += _value;

        emit Transfer(msg.sender, _to, _value);

        return true;
    }

    function approve(address _spender, uint256 _value) public returns(bool success) {
        allowance[msg.sender][_spender] = _value;
        emit Approve(msg.sender, _spender, _value);
        return true;
    }

    function getBalance(address _getaddress) public view returns (uint256 balance) {
        balance = balanceOf[_getaddress];
        return balance;
        
    }
}