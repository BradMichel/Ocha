package biud
import(
  //"gopkg.in/mgo.v2"
  //"gopkg.in/mgo.v2/bson"
)

//  Channels
type (
  Content struct{

	}




//  Message



//  User

  User struct{
    Local struct{
      UserName string
      Password string
      Email string
    }
    facebook struct{
      Id string
      UserName string
      Token string
      Email string
    }
  }
)
