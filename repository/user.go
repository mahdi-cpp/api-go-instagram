package repository

import (
	"github.com/mahdi-cpp/api-go-instagram/config"
	"github.com/mahdi-cpp/api-go-instagram/model"
)

func InitUser() {

	//var count int64 = 0
	//config.DB.Find(&model.User{}).Count(&count)
	//if count > 0 {
	//	return
	//}
	//0xba2ae424d960c26247dd6c32edc70b295c744c43

	config.DB.Create(&model.User{Username: "mahdiabdolmaleki", Phone: "09355512619", Avatar: "2018-10-23_13-55-58_UTC_profile_pic.jpg", Biography: "go lang programmer"})
	config.DB.Create(&model.User{Username: "sara_nasiry", Phone: "09207777777", Avatar: "2019-03-09_07-16-50_UTC_profile_pic.jpg"})
	config.DB.Create(&model.User{Username: "dr.abbasi.bitcoin", Phone: "09123366547", Avatar: "dr.abbasi.bitcoin_210419341_209018110984514_844281363462777318_n.jpg"})
	config.DB.Create(&model.User{Username: "halakoei", Phone: "04512121233", Avatar: "2020-01-28_01-15-41_UTC_profile_pic.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "shirin_moghaddam", Avatar: "shirin_moghaddam_121026354_4812427962108544_7736942845657074391_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "durov", Avatar: "durov_17265790_427045080983585_8337130549214707712_a.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "mamallicaa", Avatar: "mamallicaa_119119659_871855866678729_5530941536407601719_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "abednaseri", Avatar: "abednaseri_186760853_942575583255229_2875100642476586962_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "melliiic", Avatar: "melliiic_129737494_713530459592840_250000645243444580_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "whytanzania", Avatar: "whytanzania_117306842_320983259098415_5907539848736097098_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "reza_alial", Avatar: "shahab_cheraghi_85080415_766234967235941_4071192636520660992_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "meisamca", Avatar: "meisamca_212534697_346885453492622_2753089786031568080_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "elonrmuskk", Avatar: "elonrmuskk_44329317_268583430479565_454483638147350528_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "nabauti", Avatar: "nabauti_203642708_361337965370891_7694882973209500105_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "hodarostami", Avatar: "hodarostami_20065274_1598108316919320_7075841320308178944_a.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "benyamin.marco", Avatar: "benyamin.marco_121063705_202096647969008_7001926943548629903_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "jajiga.iran", Avatar: "jajiga.iran_121729554_196071998561421_8957989644433812362_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "ijustine", Avatar: "ijustine_233504118_371360141224971_6172494776591452311_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "leonardodicaprio", Avatar: "leonardodicaprio_12558345_1659293120975484_1074689227_a.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "istanbul.yari", Avatar: "istanbul.yari_245495157_378489264056990_6911603353153368661_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "rezagolzar", Avatar: "rezagolzar_208892200_277224480851626_7130548249840214103_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "zuck", Avatar: "zuck_177219615_1728341124004802_3178671336629535217_n.jpg", IsVerified: true})
	config.DB.Create(&model.User{Username: "sannevloet", Avatar: "sannevloet_217826760_326820125820581_9153130632646015297_n.jpg", IsVerified: true})

	//config.DB.Create(&model.Follow{UserID: 1, FollowerID: 3})
	//config.DB.Create(&model.Follow{UserID: 1, FollowerID: 5})
	//config.DB.Create(&model.Follow{UserID: 1, FollowerID: 6})
	//
	//config.DB.Create(&model.Follow{UserID: 3, FollowerID: 1})
	//config.DB.Create(&model.Follow{UserID: 6, FollowerID: 1})

}

func UserByID(ID uint) model.User {
	var user model.User
	config.DB.Debug().Where("id = ?", ID).Find(&user)
	return user
}

func GetUsers() []model.User {
	var users []model.User
	config.DB.Debug().Order("id ASC").Find(&users)
	return users
}

func UpdateUser() {
	config.DB.Model(model.User{}).Where("id = ?", "3").Updates(model.User{Username: "Sara Farahmand", Phone: "09125640293"})
}

func GetUser() []model.Follower {
	var currentUser = 1
	var followers []model.Follower

	config.DB.Debug().
		Table("app.follows").
		Preload("User").
		Where("app.follows.user_id = ?", currentUser).
		Order("id ASC").
		Find(&followers)

	return followers
}

func GetFollowers() []model.Follower {
	var currentUser = 1
	var followers []model.Follower

	config.DB.Debug().
		Table("app.follows").
		Preload("User").
		Where("app.follows.user_id = ?", currentUser).
		Order("id ASC").
		Find(&followers)

	return followers
}

func GetFollowing() []model.Following {
	var currentUser = 1
	var following []model.Following

	config.DB.Debug().
		Table("app.follows").
		Preload("User").
		Where("app.follows.follower_id = ?", currentUser).
		Order("id ASC").
		Find(&following)

	return following
}
