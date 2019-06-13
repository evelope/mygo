package main

import (
    "tanchi/Clib"
    "fmt"
    "math/rand"
    "os"
    "time"
)

//全局常量 界面大小
const WIDE int = 20
const HIGH int = 20

//存储食物
var food Food

//分数
var score = 0

//初始化父类 坐标
type Position struct {
    X int
    Y int
}

type Food struct {
    Position
}

//随机食物
func RandomFood() {
    food.X = rand.Intn(WIDE) + 1
    food.Y = rand.Intn(HIGH)
    //显示食物位置
    ShowUI(food.X, food.Y, '#')
}
func MapInit() {
    fmt.Fprintf(os.Stderr,
        `
  #-----------------------------------------#
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  #-----------------------------------------#
`)
}
func ShowUI(X int, Y int, ch byte) {
    //调用c语言代码绘制光标
    Clib.GotoPostion(X*2+2, Y+2)
    //将字符绘制在ui中
    fmt.Fprintf(os.Stderr, "%c", ch)
}
func PrintScore() {
    //打印分数
    Clib.GotoPostion(0, 23)
    fmt.Fprintln(os.Stderr, score)
    time.Sleep(time.Second * 2)
}

type Snake struct {
    size int
    dir  int
    pos  [WIDE * HIGH]Position
}

//初始化蛇信息
func (s *Snake) SnakeInit() {

    //初始化地图
    MapInit()
    //随机食物
    RandomFood()
    //蛇的长度
    s.size = 2
    s.pos[0].X = WIDE / 2
    s.pos[0].Y = HIGH / 2
    s.pos[1].Y = WIDE/2 - 1
    s.pos[1].Y = HIGH / 2
    //用U上L坐R右D下
    s.dir = 'R'
    for i := 0; i < s.size; i++ {
        var ch byte
        //区分蛇头和身体
        if i == 0 {
            ch = '@'
        } else {
            ch = '*'
        }
        ShowUI(s.pos[i].X, s.pos[i].Y, ch)
    }
    //go 添加一个独立函数
    //接收键盘的信息
    go func() {
        for {
            switch Clib.Direction() {
            case 72, 87, 119:
                if s.dir != 'D' {
                    s.dir = 'U'
                }
            case 80, 83, 115:
                if s.dir != 'U' {
                    s.dir = 'D'
                }
            case 65, 97, 75:
                if s.dir != 'R' {
                    s.dir = 'L'
                }
            case 68, 77, 100:
                if s.dir != 'L' {
                    s.dir = 'R'
                }
            case 32:
                s.dir = 'P'
            }
        }
    }()
}
func (s *Snake) PlayGame() {
    var dx, dy int = 0, 0
    //游戏 的流程控制
    for {
    FLAG:
        //延迟执行333s
        time.Sleep(time.Second / 3)
        //更新蛇的位置
        if s.dir == 'P' {
            goto FLAG
        }
        switch s.dir {
        case 'U':
            dx = 0
            dy = -1
        case 'D':
            dx = 0
            dy = 1
        case 'L':
            dx = -1
            dy = 0
        case 'R':
            dx = 1
            dy = 0
        }
        //蛇头和墙的碰撞
        if s.pos[0].X < 1 || s.pos[0].X >= WIDE+1 || s.pos[0].Y < 0 || s.pos[0].Y >= HIGH {
            return
        }
        //蛇头和身体的碰撞
        for i := 1; i < s.size; i++ {
            if s.pos[0].X == s.pos[i].X && s.pos[0].Y == s.pos[i].Y {
                return
            }
        }
        //蛇头的食物的碰撞
        if s.pos[0].X == food.X && s.pos[0].Y == food.Y {
            s.size++
            RandomFood()
            score++
        }
        //记录尾巴坐标
        lx := s.pos[s.size-1].X
        ly := s.pos[s.size-1].Y
        //更新蛇的坐标 蛇身体的坐标
        for i := s.size - 1; i > 0; i-- {
            s.pos[i].X = s.pos[i-1].X
            s.pos[i].Y = s.pos[i-1].Y
        }
        //蛇头的坐标
        s.pos[0].X += dx
        s.pos[0].Y += dy
        //绘制蛇的UI
        for i := 0; i < s.size; i++ {
            var ch byte
            //区分蛇头和身体
            if i == 0 {
                ch = '@'
            } else {
                ch = '*'
            }
            ShowUI(s.pos[i].X, s.pos[i].Y, ch)
        }
        //每次写完将尾巴制空
        ShowUI(lx, ly, ' ')
    }
}
func main() {
    //设置随机数种子 用作于混淆
    rand.Seed(time.Now().UnixNano())
    //隐藏控制台光标
    Clib.HideCursor()
    //创建蛇的对象
    var s Snake
    //蛇的初始化
    s.SnakeInit()
    s.PlayGame()
    PrintScore()
}