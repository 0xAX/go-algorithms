package list

import "testing"

func Test_New(t *testing.T) {
    list := New()
    list = Insert(1, list)
    list = Insert(2, list)
    list = Insert(3, list)
    list = Insert(10, list)
    list = Insert(103, list)
    list = Insert(56, list)

    has := Has(103, list)
        if has != true {
                t.Error("[Error] Has(103, list) doesn't work as expected'")
        }

        if Length(list) != 6 {
                t.Error("[Error] Length(list) doensn't work as expected'")
        }

        list = Remove(10, list)

        has = Has(10, list)
        if has != false {
                t.Error("[Error] Has(10, list) doesn't work as expected after removing 10 from list'")
        }

        if Length(list) != 5 {
                t.Error("[Error] Length(list) doensn't work as expected after removing data from list'")
        }
}
