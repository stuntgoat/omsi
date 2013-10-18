`Ordered Map: string to interface{}`
===================================

`OrderedMap` is a data structure that maps a string key to an interface{} value in Go.

Example:

    om := omsi.New()

    om.Set("cat", "funny")
    om.Set("seal", "wat")
    om.Set("dog", "happy")


    key, value, err := om.Pop()
    // dog
    fmt.Println("key", key)

    // happy
    fmt.Println("value", value)

    // nil
    fmt.Println("err", err)


    keys := om.Keys()

    // []string{'cat', 'seal'}
    fmt.Println("keys", keys)


    values := om.Values()

    // []string{'funny', 'wat'}
    fmt.Println("values", values)


    om.Delete('cat')


    values := om.Values()

    // []string{'wat'}
    fmt.Println("values", values)
