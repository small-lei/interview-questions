
# ERC-721 vs ERC-20 核心区别

## 一、设计目标差异
| 维度        | ERC-20                          | ERC-721                          |
|-------------|---------------------------------|----------------------------------|
| **代币性质** | 同质化（Fungible）              | 非同质化（Non-Fungible）         |
| **典型应用** | 加密货币（如USDT）、权益凭证    | 数字藏品、游戏道具、虚拟资产     |
| **价值特征** | 1 ETH = 1 ETH（可互换）         | 每个NFT具有唯一属性（不可互换）  |

## 二、关键接口对比
### 1. ERC-20 核心方法
```solidity
// 代币转移
function transfer(address to, uint256 amount) external returns (bool);

// 授权额度
function approve(address spender, uint256 amount) external returns (bool);

// 元数据
function name() external view returns (string memory);
function symbol() external view returns (string memory);
```

## ERC-721 特有方法
```text
// 唯一ID转移
function transferFrom(address from, address to, uint256 tokenId) external;

// 所有权查询
function ownerOf(uint256 tokenId) external view returns (address);

// 元数据扩展
function tokenURI(uint256 tokenId) external view returns (string memory);
```

## ERC-20 存储结构
```text
balances: mapping(address => uint256)  // 账户余额
allowances: mapping(address => mapping(address => uint256)) // 授权记录
```

## ERC-721 存储结构
```text
owners: mapping(uint256 => address)    // tokenId到所有者
approvals: mapping(uint256 => address) // tokenId到被授权人
tokenURIs: mapping(uint256 => string)  // 元数据链接
```

## ERC-721 陷阱
```text
// 授权攻击风险（需使用safeIncreaseAllowance）
function approve(address spender, uint256 value) public {
    require(value == 0 || allowance[msg.sender][spender] == 0);
    allowance[msg.sender][spender] = value;
}
```

## ERC-721 陷阱
```text
// 必须实现的销毁检查
function _burn(uint256 tokenId) internal virtual {
    address owner = ownerOf(tokenId);
    require(owner != address(0), "Nonexistent token");
    
    _beforeTokenTransfer(owner, address(0), tokenId);
    
    delete owners[tokenId];
    delete tokenURIs[tokenId];
}
```

## 典型应用场景
```text
ERC-20 用例
1. 稳定币(DAI/USDC)
2. 治理代币(UNI/AAVE)
3. 质押凭证

ERC-721 用例
1. CryptoPunks/CryptoKitties
2. 游戏装备(Axie Infinity)
3. 虚拟土地(Decentraland)
```









