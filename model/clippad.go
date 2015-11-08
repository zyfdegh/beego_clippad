package model

type Clippad interface{
  Clippad *model.Clippad
  New()
  Update()
  Delete()
}

type Clippad struct{
  Id string `json:id`
  OwnerId string `json:ower_id`
  Title string `json:title`
  Content string `json:content`
  TimeCreate string `json:time_create`
  TimeUpdate string `json:time_update`
}

/*
func (c *Clippad)New(){
  return nil
}

func (c *Clippad)Update(){
}

func (c *Clippad)Delete(){

}
*/
