package hagoniwa_island

import (
	"fmt"
)

// GZIP true: GZIP 圧縮転送を使用  false: 使用しない
const GZIP = false

// DEBUG true: デバッグ false: 通常
const DEBUG = false

type Init struct {
	// 各種設定値
	// プログラムを置くディレクトリ
	baseDir string
	// 画像を置くディレクトリ
	imgDir string
	// CSSファイルを置くディレクトリ
	cssDir string
	// CSSリスト
	cssList []string
	//パスワードの暗号化 true: 暗号化、false: 暗号化しない
	cryptOn bool
	// マスターパスワード
	masterPassword  string
	specialPassword string
	// データディレクトリの名前
	dirName string
	// ゲームタイトル
	title      string
	adminName  string
	adminEmail string
	urlBbs     string
	urlTopPage string
	// ディレクトリ作成時のパーミション
	dirMode int
	// 1ターンが何秒か
	unitTime int
	// 島の最大数
	maxIsland int
	// 資金表示モード
	moneyMode bool
	// トップページに表示するログのターン数
	logTopTurn int
	// ログファイル保持ターン数
	logMax int
	// バックアップを何ターンおきに取るか
	backupTurn int
	// バックアップを何回分残すか
	backupTimes int
	// 発見ログ保持行数
	historyMax int
	// 放棄コマンド自動入力ターン数
	giveupTurn int
	// コマンド入力限界数
	commandMax int
	// ローカル掲示板行数を使用するかどうか(false:使用しない、true:使用する)
	useBbs bool
	// ローカル掲示板行数
	lbbsMax int
	// 島の大きさ
	islandSize int
	// 初期資金
	initialMoney int
	// 初期食料
	initialFood int
	// 資金最大値
	maxMoney int
	// 食料最大値
	maxFood int
	// 人口の単位
	unitPop string
	// 食料の単位
	unitFood string
	// 広さの単位
	unitArea string
	// 木の数の単位
	unitTree string
	// お金の単位
	unitMoney string
	// 木の単位当たりの売値
	treeValue int
	// 名前変更のコスト
	costChangeName int
	// 人口1単位あたりの食料消費料
	eatenFood float32
	// 油田の収入
	oilMoney int
	// 油田の枯渇確率
	oilRatio int
	// ミサイル基地
	// 経験値の最大値
	maxExpPoint int // ただし、最大でも255まで
	// レベルの最大値
	maxBaseLevel  int // ミサイル基地
	maxSBaseLevel int // 海底基地
	// 経験値がいくつでレベルアップか
	baseLevelUp  []int // ミサイル基地
	sBaseLevelUp []int // 海底基地

	dBaseAuto     int
	targetIsland  int
	disEarthquake int

	// 津波
	disTsunami int
	// 台風
	disTyphoon int
	// 隕石
	disMeteo int
	// 巨大隕石
	disHugeMeteo int
	// 噴火
	disEruption int
	// 火災
	disFire int
	// 埋蔵金
	disMaizo int
	// 地盤沈下:安全限界の広さ(Hex数)
	disFallBorder int
	// 地盤沈下:その広さを超えた場合の確率
	disFalldown int
	// 怪獣:人口基準1(怪獣レベル1)
	disMonsBorder1 int
	// 怪獣:人口基準2(怪獣レベル2)
	disMonsBorder2 int
	// 怪獣:人口基準3(怪獣レベル3)
	disMonsBorder3 int
	// 怪獣:単位面積あたりの出現率(0.01%単位)
	disMonster int
	// サンジラまで
	monsterLevel1 int
	// いのらゴーストまで
	monsterLevel2 int
	// キングいのらまで(全部)
	monsterLevel3 int
	// 怪獣の種類
	monsterNumber int
	// 怪獣の名前
	monsterName []string
	// 怪獣の画像
	monsterImage []string
	// 画像ファイルその2(硬化中)
	monsterImage2 []string
	// 最低体力
	monsterBHP []int
	// 体力の幅
	monsterDHP []int
	// 特殊能力
	monsterSpecial []int
	// 経験値
	monsterExp []int
	// 死体の値段
	monsterValue []int

	// ターン杯を何ターン毎に出すか
	turnPrizeUnit int

	// 賞の名前
	prizeName []string

	// 記念碑
	monumentNumber int
	monumentName   []string

	// 画像ファイル
	monumentImage []string

	/********************
	    外見関係
	********************/
	// 大きい文字
	tagBig_ string
	_tagBig string

	// 島の名前など
	tagName_ string
	_tagName string

	// 薄くなった島の名前
	tagName2_ string
	_tagName2 string

	// 順位の番号など
	tagNumber_ string
	_tagNumber string

	// 順位表における見だし
	tagTH_ string
	_tagTH string

	// 開発計画の名前
	tagComName_ string
	_tagComName string

	// 災害
	tagDisaster_ string
	_tagDisaster string

	// ローカル掲示板、観光者の書いた文字
	tagLbbsSS_ string
	_tagLbbsSS string

	// ローカル掲示板、島主の書いた文字
	tagLbbsOW_ string
	_tagLbbsOW string

	// 順位表、セルの属性
	bgTitleCell   string // 順位表見出し
	bgNumberCell  string // 順位表順位
	bgNameCell    string // 順位表島の名前
	bgInfoCell    string // 順位表島の情報
	bgCommentCell string // 順位表コメント欄
	bgInputCell   string // 開発計画フォーム
	bgMapCell     string // 開発計画地図
	bgCommandCell string // 開発計画入力済み計画

	/********************
	    地形番号
	********************/
	landSea      int // 海
	landWaste    int // 荒地
	landPlains   int // 平地
	landTown     int // 町系
	landForest   int // 森
	landFarm     int // 農場
	landFactory  int // 工場
	landBase     int // ミサイル基地
	landDefence  int // 防衛施設
	landMountain int // 山
	landMonster  int // 怪獣
	landSbase    int // 海底基地
	landOil      int // 海底油田
	landMonument int // 記念碑
	landHaribote int // ハリボテ

	// コマンド分割
	// このコマンド分割だけは、自動入力系のコマンドは設定しないで下さい。
	// 注意：スペースは入れないように
	// ○→	'開発,0,10',  # 計画番号00〜10
	// ×→	'開発, 0  ,10  ',  # 計画番号00〜10
	commandDivido []string // 開発,0,10,  計画番号00〜10

	// 建設,11,30, // 計画番号11〜30
	// 攻撃,31,40, // 計画番号31〜40
	// 運営,41,60  // 計画番号41〜60
	commandTotal int // コマンドの種類

	// 順序
	comList []int

	// 整地系
	comPrepare     int // 整地
	comPrepare2    int // 地ならし
	comReclaim     int // 埋め立て
	comDestroy     int // 掘削
	comSellTree    int // 伐採
	comPlant       int // 植林
	comFarm        int // 農場整備
	comFactory     int // 工場建設
	comMountain    int // 採掘場整備
	comBase        int // ミサイル基地建設
	comDbase       int // 防衛施設建設
	comSbase       int // 海底基地建設
	comMonument    int // 記念碑建造
	comHaribote    int // ハリボテ設置
	comMissileNM   int // ミサイル発射
	comMissilePP   int // PPミサイル発射
	comMissileST   int // STミサイル発射
	comMissileLD   int // 陸地破壊弾発射
	comSendMonster int // 怪獣派遣
	comDoNothing   int // 資金繰り
	comSell        int // 食料輸出
	comMoney       int // 資金援助
	comFood        int // 食料援助
	comPropaganda  int // 誘致活動
	comGiveup      int // 島の放棄

	// 自動入力系
	comAutoPrepare  int // フル整地
	comAutoPrepare2 int // フル地ならし
	comAutoDelete   int // 全コマンド消去
	comName         map[int]string
	comCost         map[int]int

	// 島の座標数
	pointNumber int

	// 周围2ヘックスの座標
	ax []int // 0, 1, 1, 1, 0,-1, 0, 1, 2, 2, 2, 1, 0,-1,-1,-2,-1,-1, 0
	ay []int // 0,-1, 0, 1, 1, 0,-1,-2,-1, 0, 1, 2, 2, 2, 1, 0,-1,-2,-2

	// コメントなどに、予\定のように\が勝手に追加される
	stripslashes bool
}

// TODO: URLやパスワードなどは後に環境変数に切り出すようにする
func NewInit() *Init {
	return &Init{
		baseDir:         "http://127.0.0.1",
		imgDir:          "http://127.0.0.1/hagoniwa_img",
		cssDir:          "http://127.0.0.1/css",
		cssList:         []string{"SkyBlue.css", "Autumn.css"},
		cryptOn:         true,
		masterPassword:  "hogehoge",
		specialPassword: "hogehoge",
		dirName:         "data",
		title:           "箱庭諸島2",
		adminName:       "xxx",
		adminEmail:      "xxx@xxx.ne.jp",
		urlBbs:          "http://xxx.xxx.xxx/",
		urlTopPage:      "http://xxx.xxx.xxx/",
		dirMode:         0755,
		unitTime:        21600, // 6時間
		maxIsland:       20,
		moneyMode:       true, // true: 100の位で四捨五入, false: そのまま
		logTopTurn:      1,
		logMax:          8,
		backupTurn:      12,
		backupTimes:     0,
		historyMax:      10,
		giveupTurn:      28,
		commandMax:      20,
		useBbs:          true,
		lbbsMax:         5,
		islandSize:      12,
		initialMoney:    100,
		initialFood:     100,
		maxMoney:        9999,
		maxFood:         9999,
		unitPop:         "00人",
		unitFood:        "00トン",
		unitArea:        "00万坪",
		unitTree:        "00本",
		unitMoney:       "億円",
		treeValue:       5,
		costChangeName:  500,
		eatenFood:       0.2,
		oilMoney:        1000,
		oilRatio:        40,
		maxExpPoint:     200,
		maxBaseLevel:    5,
		maxSBaseLevel:   3,
		baseLevelUp:     []int{20, 60, 120, 200},
		sBaseLevelUp:    []int{50, 200},
		dBaseAuto:       1,
		targetIsland:    1,
		disEarthquake:   5,
		disTsunami:      15,
		disTyphoon:      20,
		disMeteo:        15,
		disHugeMeteo:    5,
		disEruption:     10,
		disFire:         10,
		disMaizo:        10,
		disFallBorder:   90,
		disFalldown:     30,
		disMonsBorder1:  1000,
		disMonsBorder2:  2500,
		disMonsBorder3:  4000,
		disMonster:      3,
		monsterLevel1:   2,
		monsterLevel2:   5,
		monsterLevel3:   7,
		monsterNumber:   8,
		monsterName: []string{
			"メカいのら",
			"いのら",
			"サンジラ",
			"レッドいのら",
			"ダークいのら",
			"いのらゴースト",
			"クジラ",
			"キングいのら",
		},
		monsterImage: []string{
			"monster7.gif",
			"monster0.gif",
			"monster5.gif",
			"monster1.gif",
			"monster2.gif",
			"monster8.gif",
			"monster6.gif",
			"monster3.gif",
		},
		monsterImage2: []string{
			"",
			"",
			"monster4.gif",
			"",
			"",
			"",
			"monster4.gif",
			"",
		},
		monsterBHP:     []int{2, 1, 1, 3, 2, 1, 4, 5, 3, 2, 3, 5, 6},
		monsterDHP:     []int{0, 2, 2, 2, 2, 0, 2, 2, 1, 2, 2, 2, 3},
		monsterSpecial: []int{0, 0, 3, 6, 1, 2, 4, 7, 5, 1, 2, 9, 10},
		monsterExp:     []int{5, 5, 7, 12, 15, 10, 20, 30, 20, 15, 20, 50, 100},
		monsterValue:   []int{0, 400, 500, 1000, 800, 300, 1500, 2000, 500, 600, 800, 3000, 3500},
		turnPrizeUnit:  100,
		prizeName: []string{
			"ターン杯",
			"繁栄賞",
			"超繁栄賞",
			"究極繁栄賞",
			"平和賞",
			"超平和賞",
			"究極平和賞",
			"災難賞",
			"超災難賞",
			"究極災難賞",
		},
		monumentNumber: 14,
		monumentName: []string{
			"モノリス",
			"平和記念碑",
			"戦いの碑",
			"トロ",
			"凱旋門",
			"記念樹",
			"慰霊碑",
			"切り株",
			"サボテン",
			"噴水",
			"斜塔",
			"時計台",
			"クリスマスツリー",
			"クリスマスツリー",
		},
		monumentImage: []string{
			"monument0.gif",
			"monument2.gif",
			"monument3.gif",
			"monument5.gif",
			"monument6.gif",
			"",
			"monument7.gif",
			"stump.gif",
			"cactus.gif",
			"fountain.bif",
			"shato.gif",
			"tokei.gif",
			"xtree1.gif",
			"xtree2.gif",
		},

		tagBig_:       `<span class="big">`,
		_tagBig:       `</span>`,
		tagName_:      `<span class="islName">`,
		_tagName:      `</span>`,
		tagName2_:     `<span class="islName2">`,
		_tagName2:     `</span>`,
		tagNumber_:    `<span class="number">`,
		_tagNumber:    `</span>`,
		tagTH_:        `<span class="head">`,
		_tagTH:        `</span>`,
		tagComName_:   `<span class="command">`,
		_tagComName:   `</span>`,
		tagDisaster_:  `<span class="disaster">`,
		_tagDisaster:  `</span>`,
		tagLbbsSS_:    `<span class="lbbsSS">`,
		_tagLbbsSS:    `</span>`,
		tagLbbsOW_:    `<span class="lbbsOW">`,
		_tagLbbsOW:    `</span>`,
		bgTitleCell:   `class="TitleCell"`,
		bgNumberCell:  `class="NumberCell"`,
		bgNameCell:    `class="NameCell"`,
		bgInfoCell:    `class="InfoCell"`,
		bgCommentCell: `class="CommentCell"`,
		bgInputCell:   `class="InputCell"`,
		bgMapCell:     `class="MapCell"`,
		bgCommandCell: `class="CommandCell"`,

		// 地形番号
		landSea:      0,  // 海
		landWaste:    1,  // 荒地
		landPlains:   2,  // 平地
		landTown:     3,  // 町系
		landForest:   4,  // 森
		landFarm:     5,  // 農場
		landFactory:  6,  // 工場
		landBase:     7,  // ミサイル基地
		landDefence:  8,  // 防衛施設
		landMountain: 9,  // 山
		landMonster:  10, // 怪獣
		landSbase:    11, // 海底基地
		landOil:      12, // 海底油田
		landMonument: 13, // 記念碑
		landHaribote: 14, // ハリボテ

		// コマンド分割
		commandDivido: []string{"開発,0,10", "建設,11,30", "攻撃,31,40", "運営,41,60"},

		// 整地系
		comPrepare:  1,
		comPrepare2: 2,
		comReclaim:  3,
		comDestroy:  4,
		comSellTree: 5,

		// 作る系
		comPlant:    11,
		comFarm:     12,
		comFactory:  13,
		comMountain: 14,
		comBase:     15,
		comDbase:    16,
		comSbase:    17,
		comMonument: 18,
		comHaribote: 19,

		// 発射系
		comMissileNM:   31,
		comMissilePP:   32,
		comMissileST:   33,
		comMissileLD:   34,
		comSendMonster: 35,

		// 運営系
		comDoNothing:  41,
		comSell:       42,
		comMoney:      43,
		comFood:       44,
		comPropaganda: 45,
		comGiveup:     46,

		// 自動入力系
		comAutoPrepare:  61,
		comAutoPrepare2: 62,
		comAutoDelete:   63,

		// 島の座標数
		pointNumber: 0,

		// 周囲2ヘックスの座標
		ax: []int{0, 1, 1, 1, 0, -1, 0, 1, 2, 2, 2, 1, 0, -1, -1, -2, -1, -1, 0},
		ay: []int{0, -1, 0, 1, 1, 0, -1, -2, -1, 0, 1, 2, 2, 2, 1, 0, -1, -2, -2},

		// 文字列の削除
		stripslashes: false,
	}
}

func (i *Init) setVariable() {
	i.pointNumber = i.islandSize * i.islandSize

	i.comList = []int{
		i.comPrepare,
		i.comSell,
		i.comPrepare2,
		i.comReclaim,
		i.comDestroy,
		i.comSellTree,
		i.comPlant,
		i.comFarm,
		i.comFactory,
		i.comMountain,
		i.comBase,
		i.comDbase,
		i.comSbase,
		i.comMonument,
		i.comHaribote,
		i.comMissileNM,
		i.comMissilePP,
		i.comMissileST,
		i.comMissileLD,
		i.comSendMonster,
		i.comDoNothing,
		i.comMoney,
		i.comFood,
		i.comPropaganda,
		i.comGiveup,
		i.comAutoPrepare,
		i.comAutoPrepare2,
		i.comAutoDelete,
	}

	// 計画の名前と値段
	i.comName = map[int]string{
		i.comPrepare:      "整地",
		i.comPrepare2:     "地ならし",
		i.comReclaim:      "埋め立て",
		i.comDestroy:      "掘削",
		i.comSellTree:     "伐採",
		i.comPlant:        "植林",
		i.comFarm:         "農場整備",
		i.comFactory:      "工場建設",
		i.comMountain:     "採掘場整備",
		i.comBase:         "ミサイル基地建設",
		i.comDbase:        "防衛施設建設",
		i.comSbase:        "海底基地建設",
		i.comMonument:     "記念碑建造",
		i.comHaribote:     "ハリボテ設置",
		i.comMissileNM:    "ミサイル発射",
		i.comMissilePP:    "PPミサイル発射",
		i.comMissileST:    "STミサイル発射",
		i.comMissileLD:    "陸地破壊弾発射",
		i.comSendMonster:  "怪獣派遣",
		i.comDoNothing:    "資金繰り",
		i.comSell:         "食料輸出",
		i.comMoney:        "資金援助",
		i.comFood:         "食料援助",
		i.comPropaganda:   "誘致活動",
		i.comGiveup:       "島の放棄",
		i.comAutoPrepare:  "整地自動入力",
		i.comAutoPrepare2: "地ならし自動入力",
		i.comAutoDelete:   "全計画を白紙撤回",
	}

	i.comCost = map[int]int{
		i.comPrepare:  5,
		i.comPrepare2: 100,
		i.comReclaim:  150,
		i.comDestroy:  200,
		i.comSellTree: 0,

		i.comPlant:        20,
		i.comFarm:         300,
		i.comFactory:      300,
		i.comMountain:     500,
		i.comBase:         500,
		i.comDbase:        1000,
		i.comSbase:        2000,
		i.comMonument:     9999,
		i.comHaribote:     10,
		i.comMissileNM:    0,
		i.comMissilePP:    0,
		i.comMissileST:    0,
		i.comMissileLD:    0,
		i.comSendMonster:  10000,
		i.comDoNothing:    0,
		i.comSell:         0,
		i.comMoney:        0,
		i.comFood:         0,
		i.comPropaganda:   1000,
		i.comGiveup:       0,
		i.comAutoPrepare:  0,
		i.comAutoPrepare2: 0,
		i.comAutoDelete:   0,
	}
}

func main() {
	init := NewInit()
	fmt.Println(init.baseDir)
	fmt.Println(init.cssList)
	fmt.Println(init.cryptOn)

	// 日本時間にあわせる
	// 海外のサーバに設置する場合は次の行にある//をはずす。
	// os.Setenv("TZ", "JST-9")

	init.setVariable()
}
