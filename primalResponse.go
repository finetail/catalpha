package catalpha

import ("gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type primal struct{
	Id bson.ObjectId `bson:"_id"`
	Name string `bson:"name"`
	Element string `bson:"element"`
	Bless string `bson:"bless"`
	Bless3 string `bson:"bless3"`
	Bless4 string `bson:"bless4"`
	BlessSub string `bson:"blessSub"`
	BlessSub3 string `bson:"blessSub3"`
	BlessSub4 string `bson:"blessSub4"`
	MinHP string `bson:"MinHP"`
	MinAtk string `bson:"MinAtk"`
	MaxHP100 string `bson:"MaxHP100"`
	MaxAtk100 string `bson:"MaxAtk100"`
	MaxHP150 string `bson:"MaxHP150"`
	MaxAtk150 string `bson:"MaxAtk150"`
}

func responsePrimal(in string) (out string){

	dbsession, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	defer dbsession.Close()

	primals := dbsession.DB("graburu").C("primal")
	result := primal{}
	fmt.Printf("%s",in)
	query := bson.M{"name":in}
	//query := bson.M{}
	primals.Find(query).One(&result) //何故か取れないんですけど！！！！
	return fmt.Sprintf("name：%s bless3：%s",result.Name,result.Bless3)
}