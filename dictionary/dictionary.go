package dictionary

import (
	"math/rand"
	"time"
)

// 黑话词典

var (
	// 两字动词
	V2 = []string{
		"化反", "皮实", "复盘", "赋能", "加持", "沉淀", "倒逼", "落地", "串联", "协同", "反哺", "兼容", "包装",
		"重组", "履约", "响应", "量化", "发力", "布局", "联动", "细分", "梳理", "输出", "加速", "共建", "共创",
		"支撑", "融合", "解耦", "聚合", "集成", "对标", "对齐", "聚焦", "抓手", "拆解", "拉通", "抽象", "摸索",
		"提炼", "打通", "吃透", "迁移", "分发", "分层", "封装", "辐射", "围绕", "复用", "渗透", "扩展", "开拓",
		"给到", "死磕", "破圈", "拉齐", "迭代"}

	// 两字名词
	N2 = []string{
		"漏斗", "中台", "闭环", "生态", "打发", "补位", "体系", "纽带", "矩阵", "刺激", "规模", "场景", "维度",
		"格局", "形态", "生态", "话术", "认知", "玩法", "体感", "感知", "调性", "心智", "战役", "合力", "赛道",
		"基因", "因子", "模型", "载体", "横向", "通道", "链路", "试点", "用户", "痛点", "痒点", "爽点", "线上",
		"头部", "腰部"}

	// 三字名词
	N3 = []string{
		"新生态", "感知度", "颗粒度", "方法论", "组合拳", "引爆点", "点线面", "精细化", "差异化", "平台化", "结构化",
		"影响力", "耦合性", "易用性", "便捷性", "一致性", "端到端", "点到点", "短平快", "护城河", "新次元", "UGC"}

	// 四字短语
	P4 = []string{"持久收益", "底层逻辑", "产品赋能", "顶层设计", "交付价值", "生命周期", "价值转化", "强化认知",
		"资源倾斜", "完善逻辑", "抽离透传", "复用打法", "评判标准", "商业模式", "快速响应", "定性定量", "关键路径",
		"去中心化", "结果导向", "垂直领域", "归因分析", "体验度量", "信息屏障", "价值回归", "如何收口", "用户画像",
		"深度共建", "自然势能", "用户感知", "价值链路", "用户体感", "流量风口", "裂变增长", "原子表达", "语义隔膜",
		"延迟满足", "降维打击"}

	// 五字短语
	P5 = []string{"技术栈下移"}
)

// 返回随机动词
func GetRandomV() string {
	rand.Seed(time.Now().UnixNano())
	return V2[rand.Intn(len(V2))]
}

// 返回随机名词
func GetRandomN() string {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(1)%2 == 0 {
		return N3[rand.Intn(len(N3))]
	}
	return N2[rand.Intn(len(N2))]
}

// 返回随机短语
func GetRandomP() string {
	rand.Seed(time.Now().UnixNano())
	return P4[rand.Intn(len(P4))]
}
