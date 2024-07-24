package global

const (
	StdOut   = "./logs/info.log"
	StdError = "./logs/error.log"
	DebugError = "./logs/debug.log"
	LoginRoleSuper = 0
	LoginRoleCompany = 1
	LoginRoleShop = 2
	ExpressStore     = 1 //门店自提
	ExpressLocal     = 2 //同城配送
	ExpressEms = 3 //物流配送

	//商品目录
	GoodsPath       = "goods"
	SysName         = "动创云订货配送"
	Describe        = "致力于解决订货渠道"
	RoleSuper       = 80 //超管
	RoleAdmin     = 81 //子管理员
	RoleUser = 82 //用户
	RoleSaleMan = 85 //业务员
	RoleDriver = 86 //司机
	RegisterUserVerify = 1 //新用户需要审核,通过后才可以登录
	RegisterUserLogin = 2 //新用户直接注册+登录
	//用户关闭的
	SysUserDisable = 1
	//用户是开启的
	SysUserSuccess = 2
	//订单状态
	OrderStatusPayFail = -2 //支付失败

	OrderStatusClose = -1 //订单关闭

	OrderStatusWaitPay = 0 //默认状态，就是待支付

	OrderStatusWaitSend = 1 //备货中

	OrderWaitConfirm = 2 //配送中/待取

	OrderWaitRefunding = 3 //售后处理中

	OrderStatusCancel = 4 //大B操作作废

	OrderStatusReturn = 5 //售后处理完毕

	OrderPayStatusOfflineWait = 6 //线下付款 未付款

	OrderPayStatusOfflineSuccess = 7 //线下收款 已收款

	OrderStatusPaySuccess = 8 //线上支付成功

	OrderStatusWaitPayDefault = 9 //下单了,待支付,会先放在redis中的

	OrderStatusOver = 10 //订单收尾,那就是收货了,确认了

	OrderApproveOk = 11 //审核通过了

	OrderApproveReject = 12 //审批驳回

	//大B资源限制
	CompanyVip           = 10   //大B最多可以设置
	CompanyLine          = 1   //默认路线
	CompanyOrderRange = 18 //默认订单查询时间范围 单位为月
	CompanyMaxRole       = 20  //大B最多可以设置个角色
	CompanyMaxGoods      = 500 //大B最多可以创建个商品
	CompanyMaxShop       = 10000  //大B最多可以创建个小B客户
	CompanyMaxGoodsClass = 50  //大B最多可以设置分类个数
	CompanyMaxGoodsTag   = 30  //大B最多可以设置标签个数
	CompanyMaxGoodsImage = 4   //大B最多可以设置单个商品做多张图片
	CompanyUserTag       = 30  //大B最多可以设置客户标签个数
	CompanySmsNumber = 100 //大B默认的可用短信条数,因为短信涉及到缴费,所以短信的可用条数是在其他表中(CompanyEmsQuotaCnf)颗粒度控制
	CompanyIndexMessage = 3 //首页消息条目
	CompanyIndexAds = 5 //广告数量
	CompanyLineBindShop = 500 //路线最多支持绑定多少个客户
	CompanyExportWorker = 5 //导出任务队列个数
	CompanySalesmanNumber = 20 //大B默认拥有个业务员
	CompanySmsRecordTag = true //大B的短信消费记录，默认是开启
	OffLinePay = 8 //大B最多可以设置线下支付的个数
	CompanyMaxLocal = 10 //大B支持最多自提设置,不支持动态配置 现在是固定的
	CompanyMaxUnit = 100 //商品单位100个
	CompanyMaxBrand = 200 //商品品牌

	OrderLayerKey    = " layer desc,id desc "
	OrderTimeKey     = "created_at desc"
	//分表的逻辑
	SplitOrder                 = 1
	SplitOrderDefaultTableName = "orders"
	//关联的订单子表,如果进行了订单表的分割,也会默认进行一个分割
	SplitOrderDefaultSubTableName = "order_specs"
	//扩展表
	SplitOrderExtendSubTableName = "order_extend"

	//周期配送下单索引表
	SplitOrderCycleSubTableName = "order_cycle_cnf"

	//订单修改表
	SplitOrderEdit = "order_edit_record"

	//订单退换货表
	SplitOrderReturn = "order_return"

	//出入库记录流水表
	InventoryRecordLog = "inventory_record"

	//Cycle 配送的设置
	//每天
	CyCleTimeDay = 1
	//每周
	CyCleTimeWeek = 2


	//售后状态
	RefundDefault = 1 //售后处理中
	RefundOk = 2 //售后处理完成 也就是同意
	RefundOkOverReject = -1 //大B驳回
	RefundOkCancel = -2 //用户主动撤销
	RefundCompanyCancelCType = 3 //大B作废操作

)

func AICompany()  {
	
}
func GetExpressCn(v int) string {
	switch v {
	case ExpressStore:
		return "门店自提"
	case ExpressLocal:
		return "周期配送"
	case ExpressEms:
		return "物流配送"
	}
	return "周期配送"
}
func GetRoleCname(v int) string  {
	switch v {
	case RoleSuper:
		return "超管"
	case RoleAdmin:
		return "子管理员"
	case RoleUser:
		return "用户"


	}
	return "非法角色"
}




