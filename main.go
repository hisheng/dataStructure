package main

func main()  {
	btree := tree{root:nil}
	btree.Insert(1)
	btree.Insert(2)
	btree.Insert(3)
	btree.Insert(4)
	btree.Insert(5)
	btree.Insert(6)
	btree.Insert(7)

	btree.breadthTravel()
}


