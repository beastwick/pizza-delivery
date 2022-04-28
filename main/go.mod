module koneska.pizza/main

go 1.18

replace koneksa.pizza/dispatcher => ../dispatcher

replace koneksa.pizza/deliverer => ../deliverer

require (
	koneksa.pizza/deliverer v0.0.0-00010101000000-000000000000
	koneksa.pizza/dispatcher v0.0.0-00010101000000-000000000000
)

require (
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/goccy/go-graphviz v0.0.9 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/ofabry/go-callvis v0.6.1 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/profile v1.6.0 // indirect
	golang.org/x/image v0.0.0-20220321031419-a8550c1d254a // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/sys v0.0.0-20220325203850-36772127a21f // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)
