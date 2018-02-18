package pairs

func Valid(pair string) bool {
for _, v := range validPairs {
if v.String() == pair {
return true
}
}
return false
}

func Find(pair string) *AssetPair {
for _, v := range validPairs {
if v.String() == pair {
return &v
}
}
return nil
}