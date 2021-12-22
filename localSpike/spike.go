/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/21/2021
    @Note:       本地库存处理函数
**/

package localSpike

type LocalSpike struct {
	LocalInStock     int64 `mapstructure:"local_in_stock" json:"local_in_stock" yaml:"local_in_stock"`
	LocalSalesVolume int64 `mapstructure:"local_sales_volume" json:"local_sales_volume" yaml:"local_sales_volume"`
}

//本地扣库存,返回bool值
func (spike *LocalSpike) LocalDeductionStock() bool {
	spike.LocalSalesVolume = spike.LocalSalesVolume + 1
	return spike.LocalSalesVolume <= spike.LocalInStock
}
