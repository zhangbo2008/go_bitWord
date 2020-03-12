////https://articles.zsxq.com/id_wr18nw85itm3.html
//
//本系列整理了10个工作量和难度适中的Golang小项目，适合已经掌握Go语法的工程师进一步熟练语法和常用库的用法。
//
//
//
//问题描述：
//
//有一组非负整数，实现一个位向量类型，能在O(1)时间内完成插入、删除和查找等操作。
//
//
//
//要点：
//
//实现Has(uint)、Add(uint)、Remove(uint)、Clear()、Copy()、String()、AddAll(…uint)、UnionWith()、IntersectWith()、DifferenceWith()、SymmetricDifference()方法。
//
//
//拓展：
//
//使用uint存储而不是uint32或uint64这样限定字长的类型。
//
//
/*
位向量的实现.  基本就是bitmap算法.简单的1b/
只不过拓展了一下就是,存的不是数字了,而是一个word,word用一个数字表示.
本质还是数字.





*/











//代码实现：
package main

import (
  "bytes"
  "fmt"
)

//IntSet   这个就是位向量. 是无符号整数组成的数组.
// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
  words []uint //注意这里面每一个uint占64位.
}


//下面写方法


//给一个整数n,返回一个数值表示这个整数的二进制码有多少个1
func (s *IntSet) countBit(n uint) int {
  count := 0

  for n != 0 {
    n = n & (n - 1)   //每一次去掉数位里面最低的1.
    count += 1// 所以有多少个1,就加多少次.
  }

  return count
}









// 输入一个整数,返回这个整数可以表示多少个单词和剩余的int
//  一个单词这里面使用64个bit来表示.

func (s *IntSet) calWordBit(x int) (word int, bit uint) {
  word, bit = x/wordSize, uint(x%wordSize)
  return
}

const wordSize = 32 << (^uint(0) >> 63)




// 就是bitmap算法.而已.
// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
  word, bit := s.calWordBit(x)
  return word < len(s.words) && s.words[word]&(1<<bit) != 0
}






// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
  word, bit := s.calWordBit(x)
  for word >= len(s.words) {// 位置不够就不停的加0,
    s.words = append(s.words, 0)
  }
  s.words[word] |= 1 << bit   //加入1.
}

func (s *IntSet) AddAll(nums ...int) {
  for _, n := range nums {
    s.Add(n)
  }
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
  for i, tword := range t.words {


    if i < len(s.words) {
      s.words[i] |= tword
    } else {
      s.words = append(s.words, tword)
    }


  }
}





// Set s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] &= tword//改成&即可.
    } else {
      s.words = append(s.words, tword)
    }
  }
}

// Set s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] &^= tword//功能同a&(^b)相同,返回的东西一定是t里面没有的,在s中有的.
    } else {
      s.words = append(s.words, tword)
    }
  }
}

// Set s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
  for i, tword := range t.words {
    if i < len(s.words) {
      s.words[i] ^= tword   //返回一定是2个中不一样的.
    } else {
      s.words = append(s.words, tword)
    }
  }
}

// return the number of elements
func (s *IntSet) Len() int {
  count := 0
  for _, word := range s.words {
    count += s.countBit(word)
  }
  return count
}

// remove x from the set
func (s *IntSet) Remove(x int) {
  word, bit := s.calWordBit(x)
  s.words[word] &^= 1 << bit
}

// remove all elements from the set
func (s *IntSet) Clear() {
  for i := range s.words {
    s.words[i] = 0
  }
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
  new := &IntSet{}
  new.words = make([]uint, len(s.words))
  copy(new.words, s.words)
  return new
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
  var buf bytes.Buffer
  buf.WriteByte('{')
  for i, word := range s.words {
    if word == 0 {
      continue
    }
    for j := 0; j < wordSize; j++ {
      if word&(1<<uint(j)) != 0 {
        if buf.Len() > len("{") {
          buf.WriteByte(' ')
        }
        fmt.Fprintf(&buf, "%d", wordSize*i+j)
      }
    }
  }
  buf.WriteByte('}')
  return buf.String()
}

// Return set elements.
func (s *IntSet) Elems() []int {
  e := make([]int, 0)
  for i, word := range s.words {
    for j := 0; j < wordSize; j++ {
      if word&(1<<uint(j)) != 0 {
        e = append(e, i*wordSize+j)
      }
    }
  }
  return e
}
