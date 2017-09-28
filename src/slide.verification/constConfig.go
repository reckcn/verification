package slide_verification

const (
	// 裁剪的小图大小
	shearSize = 40
	// 原始图片的所在路径 300*300
	icturePath = ""
	// 原始图片的数量
	imgNum = 60
	// 原始图片的宽px
	imgWidth = 300
	// 原始图片的高px
	imgHeight = 300
	// 裁剪位置X轴最小位置
	minRangeX = 30
	// 裁剪位置X轴最大位置
	maxRangeX = 240
	// 裁剪位置Y轴最小位置
	minRangeY = 30
	// 裁剪位置Y轴最大位置
	maxRangeY = 200
	// 裁剪X轴大小 裁剪成20张上10张下10张
	cutX = 30
	// 裁剪Y轴大小 裁剪成20张上10张下10张
	cutY = 150
	// 允许误差 单位像素
	deviationPx = 2
	// 是否跨域访问 在将项目做成第三方使用时可用跨域解决方案 所有的session替换成可共用的变量(Redis)
	isCallback = false
	// 最大错误次数
	maxErrorNum = 4
)

// 小图相对原图左上角的x坐标  x坐标保存到session 用于校验
var positionX int

// 小图相对原图左上角的y坐标  y坐标返回到前端
var positionY int

// action命令列表
var actionList = [3]string{"getCode", "check", "checkResult"}

// 图片规格列表 默认300*300
var imgSpecList = [3]string{"300*300", "300*200", "200*100"}
