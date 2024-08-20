package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2" // imports as package "cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "my-cli"
	app.Usage = "用urfave构建的命令行程序."

	// 子命令 list
	listCmd := cli.Command{
		Name:  "list",
		Usage: "显示文件列表",
		Action: func(c *cli.Context) error {
			//获取当前工作目录
			dir, err := os.Getwd()
			if err != nil {
				return err
			}
			// 获取当前目录下的所有文件
			files, err := os.ReadDir(dir)
			if err != nil {
				return err
			}
			// 遍历文件列表
			for _, file := range files {
				fmt.Println(file.Name())
			}
			// 返回 nil，表示执行成功
			return nil
		},
	}

	// 子命令 add
	addCmd := cli.Command{
		Name:  "add",
		Usage: "添加文件",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "filename",
				Aliases: []string{"fn"},
				Usage:   "文件名",
			},
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "强制添加文件",
			},
		},
		Action: func(c *cli.Context) error {
			filename := c.String("filename")
			force := c.Bool("force")

			// 检查文件是否已存在，如果存在并且没有使用force标志，则返回错误
			_, err := os.Stat(filename)
			if !os.IsNotExist(err) && !force {
				return fmt.Errorf("文件 %s 已存在，请使用 --force 强制覆盖", filename)
			}

			// 创建文件
			f, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer f.Close()
			fmt.Println("添加文件成功", filename)
			return nil
		},
	}

	// 添加子命令 delete
	deleteCmd := cli.Command{
		Name:  "delete",
		Usage: "删除文件",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "filename",
				Aliases: []string{"fn"},
				Usage:   "文件名",
			},
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "强制",
			},
		},
		Action: func(c *cli.Context) error {
			filename := c.String("filename")
			force := c.Bool("force")
			// 检查文件是否存在，如果不存在并且没有使用 force 标志，则返回错误
			_, err := os.Stat(filename)
			if os.IsNotExist(err) && !force {
				return fmt.Errorf("文件 %s 不存在", filename)
			}
			// 删除文件
			err = os.Remove(filename)
			if err != nil {
				return err
			}
			fmt.Println(filename, "删除成功")
			return nil

		},
	}

	app.Commands = []*cli.Command{
		&listCmd,
		&addCmd,
		&deleteCmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
