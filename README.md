# statedb2cosmosgenesis
将以太坊的statedb转换为cosmos sdk的genesis.json。
- 目的：以前使用tendermint+go-ethereum开发一条链zero，类似ethermint项目，由于开发能力有限，该链不够健壮，离公链标准还有一定差距；而evmos基于cosmos sdk，集成了evm，功能上包含了原链的功能，且社区强大，代码健壮，可维护性好；因此，希望将原链迁移到evmos来
- 不是真的将以太坊的statedb进行直接转换，而是先将statedb dump为json文件，再将json中的账户、合约导入到基于cosmos sdk开发的链的genesis.json中

## 操作步骤
- 停止zero链，导出zero_dev.genesis.json
- 初始化一条evmos链【一定记得修改evm的原生币名称为aevmos】，并启动，至少出一个块
- 导出evmos链为 genesis_export_number2.json 出块到2
- 执行合并./evmtoevmos genesis_export_number2.json zero_dev.genesis.json > genesis_compound.json
- 复制一个genesis_export_number2.json为genesis_new.json，并编辑以下内容:
- 初始化高度为需求值：zero链最后一个高度+1
- 修改bank模块的supply aevmos数量：首次执行时也会统计提示具体数额，也可以根据需求进行修改
- 启动:先unsafe-reset-all 再 start

## 注意事项
- 原链zero链使用go-ethereum v1.10.3,而evmosd v1.1.0 使用go-ethereum v1.10.16