package core

import "fmt"

func (v *Version) removeDuplicateObjects(fields map[string]*Field, currentFieldCount int, async bool) int {
	
	if v == nil {
		return currentFieldCount
	}

	for _, o := range v.Objects { //[]->[]
		//files 
		x := 0
		// for i := 0 ;i<len(o.Fields);i++ {

		// }
		//versions[] 1version
		//	objects[] 123,field1 - 123,field1 
		//	fields[]
		//////
		//versions[] 2version
		//	objects[] 123,field1 - 123,field1 
		//	fields[]
		//x = 0 
		//map(fields)  1 2 3 

		//2 1 3 1 2 (key of fields) 
		//    x
		//2nd approach 
		//0 1 2 3 4 5 6 
		//1 2 3 2 4 5   (on) time compelexity 
		//1 2 3 
		
		// res := []*Field{}
		
		for i, f := range o.Fields {
			fIDExist := true 
			fKeyExist := true 

			if  _, exists := fields[f.ID]; !exists {
				fields[f.ID] = f
				fIDExist = false 
				// res = append(res, f)
			} 
			
			if _, exists := fields[f.Key]; !exists {
				fields[f.Key] = f
				fKeyExist = false 
				// res = append(res, f)
			} 
			
			if fIDExist || fKeyExist {
				//remove it 
				if async {
					o.Lock()
				}
				o.Fields[i] = o.Fields[x]
				if async {
					o.Unlock()
				}
				x++
				currentFieldCount--
			}
		}
		if async {
			o.Lock()
		}
		o.Fields = o.Fields[x:]
		if async {
			o.Unlock()
		}
	}

	return currentFieldCount
}
// k = n(version) 
// m = nobjects 
// n = n(fields) 
// k*m*n

func (v *Version) removeDuplicateViews(views map[string]*View, currentViewCount int, async bool) int {
	if v == nil {
		return currentViewCount
	}

	for _, s := range v.Scenes {
		//views
		x := 0
		for i, vw := range s.Views {
			
			if  _, exists := views[vw.ID+vw.Key]; !exists {
				views[vw.ID+vw.Key] = vw
			} else {
				//remove it 
				if async {
					s.Lock()
				}

				s.Views[i] = s.Views[x]
				if async {
					s.Unlock()
				}
				x++
				//update viewcount 
				currentViewCount--
			}
		}
		if async {
			s.Lock()
		}
		s.Views = s.Views[x:]
		if async {
			s.Unlock()
		}
	}

	return currentViewCount
}

//return new fieldCount after removal
func (v *Version) removeDuplicateObjectsAsync(fields map[string]*Field, oldFieldCount int) chan int {
	res := make(chan int, 1)

	go func() {
		if v == nil {
			res <- oldFieldCount
		}
		newFieldCount := v.removeDuplicateObjects(fields, oldFieldCount, true)
		res <- newFieldCount
	}()

	return res 
}

func (v *Version) removeDuplicateViewsAsync(views map[string]*View, oldViewCount int) chan int {
	res := make(chan int, 1)

	go func() {
		if v == nil {
			res <- oldViewCount
		}
		newViewCount := v.removeDuplicateViews(views, oldViewCount, true)
		res <- newViewCount
	}()
	
	return res 
}

func (obj *DBSchema) RemoveDuplicates(concurrent bool) {
	
	// 50 reques per scond 
	// 5 MBs * 50 = 250 MBs
	// 250 MBs * 60 = 15 GBs

	// 
	
	
	// Uploading (25 MB) => LBs => Direct Call () (Hot calls)
	// 	--> S3 	=> SQS 
	//	Worker Nodes <- (ASG) CPU/Network
	//		 |
	//		 V
	//		LB => API => SNS 
	
	//go routines 
	if obj == nil {
		return
	}

	if len(obj.Versions) == 0 {
		return
	}

	//The description was not clear whether we do have duplicates in versions level as well, in other words, if we might have multiple versions with the same id. in such case I would do the same concept
	//I would use a map to store the ids and check if the id exists in the map, if it does, I would remove it from the array, otherwise I would add it to the map
	//In case we have so many ids, I can do this 
	
	
	for _, v := range obj.Versions  {
		//I assumed duplications are per version and we need to clean versions separately.
		// fields and views maps are used per one version (meaning I assumed duplicates happening in the version level only) 
		// 		so, in case we need to unify that to all versions, we can move the maps up outside the loop
		fields := map[string]*Field{}
		views := map[string]*View{}
		
		if concurrent {
			fmt.Println("Concurrent processing")
			chFieldCount := v.removeDuplicateObjectsAsync(fields, obj.FieldCount)
			chViewCount := v.removeDuplicateViewsAsync(views, obj.ViewCount)
			obj.FieldCount = <- chFieldCount
			obj.ViewCount = <- chViewCount
		} else {
			fmt.Println("Non-concurrent processing")
			obj.FieldCount = v.removeDuplicateObjects(fields, obj.FieldCount, false)
			obj.ViewCount = v.removeDuplicateViews(views, obj.ViewCount, false)
		}
		
		
	}

}
