package global

const (
	// IDC类型
	IdcIDC  = 1
	IDCACDN = 2
	IDCPCDN = 3

	// 机房归属

	BelongZJ = 1 //自建
	BelongZM = 2 //招募

	//计费方式

	ChargingBuyout = 1 //买断
	Charging95     = 2 //95

	IdcStatusUp      = 1 //上架
	IdcStatusLoading = 2 //待交付
	IdcStatusTest    = 3 //待测试
	IdcStatusOffline = 4 //下架
	BandwidthZX      = 1 //专线
	BandwidthNZX     = 2 //非专线
	YIDONG           = 1
	DIANXIN          = 2
	LIANTONG         = 3
	Other            = 3
	HostLoading      = 0 //链接中
	HostUp           = 1 //在线
	HostOffline      = 1 //离线//非专线

)
