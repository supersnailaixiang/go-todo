package common

func ToDoStatus(status int) string {
	switch status {
	case 0:
		return "待完成"
	case 5:
		return "已放弃"
	case 10:
		return "已完成"
	default:
		return "状态错误"
	}
}
