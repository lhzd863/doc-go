package db

import (
        "database/sql"

        _ "github.com/lib/pq"
        "fmt"
)
type Pgdb struct {
        Host     string
        Port     int
        User     string
        Password string
        Dbname   string
        Conn   *sql.DB
}

func NewConnectDB(host string, port int, usr string, passwd string, dbname string) (*Pgdb, error){
        psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
                "password=%s dbname=%s sslmode=disable",
                host, port, usr, passwd, dbname)

        db, err := sql.Open("postgres", psqlInfo)
        if err != nil {
                panic(err)
        }

        //err = db.Ping()

        return &Pgdb{
           Conn:        db,
           Host:        host,
           Port:        port,
           User:        usr,
           Password:    passwd,
           Dbname:      dbname,
        }, nil
}


func (pg *Pgdb) QueryWeather(sql string) map[int]map[string]string {
        var wr WeatherRet
        ret := make(map[int]map[string]string)
        rows,err:=pg.Conn.Query(sql)

        if err!= nil{
                fmt.Println(err)
        }
        defer rows.Close()
        i:=0
        for rows.Next(){
                err:= rows.Scan(&wr.City,&wr.Cityid,&wr.Temp,&wr.WD,&wr.WS,&wr.SD,&wr.AP,&wr.Njd,&wr.WSE,&wr.Time,&wr.Sm,&wr.IsRadar,&wr.Radar)
                if err!= nil{
                        fmt.Println(err)
                }
                tmap := make(map[string]string)
                tmap["City"] = wr.City
                tmap["Cityid"] = wr.Cityid
                tmap["Temp"] = wr.Temp
                tmap["WD"] = wr.WD
                tmap["WS"] = wr.WS
                tmap["AP"] = wr.AP
                tmap["Njd"] = wr.Njd
                tmap["WSE"] = wr.WSE
                tmap["Time"] = wr.Time
                tmap["SM"] = wr.Sm
                tmap["IsRadar"] = wr.IsRadar
                tmap["Radar"] = wr.Radar
                ret[i] = tmap
                i++
        }

        err = rows.Err()
        if err!= nil{
                fmt.Println(err)
        }
        return ret
}

type WeatherRet struct{
        City    string
        Cityid  string
        Temp    string
        WD      string
        WS      string
        SD      string
        AP      string
        Njd     string
        WSE     string
        Time    string
        Sm      string
        IsRadar string
        Radar   string
}
