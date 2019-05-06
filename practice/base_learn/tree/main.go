package main

import (
	"encoding/json"
	"fmt"
)

type tree struct {
	ID       int64   `json:"id"`
	ParentID int64   `json:"parent_id"`
	Name     string  `json:"name"`
	Children []*tree `json:"children,omitempty"`
}

func main() {

	ts := make([]*tree, 0)
	top := make(map[int64]*tree)
	res := make([]*tree, 0)

	//query db
	// rows, err := db.Query(c, xxxSQL)
	// if err != nil {
	// 	log.Error("d.db.Query error(%v)", err)
	// 	return
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	t := &tree{}
	// 	if err = rows.Scan(&t.ID, &t.ParentID, &t.Name); err != nil {
	// 		log.Error("rows.Scan error(%v)", err)
	// 		return
	// 	}
	// 	ts = append(ts, t)
	// }

	//mock data
	t1 := &tree{ID: 1, ParentID: 0, Name: "a"}
	ts = append(ts, t1)
	t2 := &tree{ID: 2, ParentID: 0, Name: "b"}
	ts = append(ts, t2)

	t11 := &tree{ID: 3, ParentID: 1, Name: "aa"}
	ts = append(ts, t11)
	t22 := &tree{ID: 4, ParentID: 2, Name: "bb"}
	ts = append(ts, t22)

	for _, t := range ts {
		if t.ParentID == 0 {
			top[t.ID] = t
			res = append(res, t)
		}
	}
	for _, t := range ts {
		a, ok := top[t.ParentID]
		if ok && a != nil {
			a.Children = append(a.Children, t)
		}
	}

	js, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(js))
}
