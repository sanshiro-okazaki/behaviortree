# behaviortree

## Usage

Go
```
$ go mod tidy && go build -o bt
$ SKILL_A=0.5 ./bt
```

Make
```
$ make run -s=0.5
```

Docker
```
$ docker build -t behaviortree .
$ docker run -e SKILL_A=0.5 behaviortree
```

## Input
スキルA発動確率
```
SKILL_A=0.5
```

## Example
```
$ make run -s=0.5
go mod tidy && go build -o bt && SKILL_A=0.5 ./bt
出発
敵に寄る
友達1
友達2
スキルB発動
スキルB発動
終了1
&main.Player{HP:150, SkillA:0.5, SkillB:1} &main.Enemy{HP:0}
```
