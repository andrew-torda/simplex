# simplex
Nelder Meade simplex optimization method.

The basic interface is simple. You ask for a new object and then say Run. There is an absurd number of options. You can specify upper and lower bounds for parameters individually, convergence requirements for each parameter and there are two different methods for setting up the initial vertices. The most complicated option invokes a second cost function on the data each time a new best point is found. If one is doing an optimisation for some kind of fitting, you can have a standard cost function that acts on most of the data, but a second function that acts on part of the data for testing (to see recall versus generalisation or overfitting).

## Introduction

There are lots of tests which give a flavour of what one should do. The example_xx_test.go files are real examples. The simplest start is to define your cost function. It takes a slice of float32's as its only argument and returns a float32 and error. A quadratic function might look like this

 function cost ( x []float32) (result, error) {
     a := x[0] - 5
     b := x[1] - 2
     return a * a + b * b, nil
 }

This would have minima at `x[0] = 5` and `x[1] = 2`. It does not allow for errors, so it returns nil as the second result. 
If you need additional arguments, you will need some kind of wrapper. Maybe you have some fixed arguments which are not to be optimized. A closure might be the best approach. There are lots of closures in the tests.

You then need to make a simplex control argument
 	s := NewSplxCtrl(cost, iniPrm, maxsteps)
where `cost` is the cost function and iniPrm is a slice with initial parameters. Maybe it was set by a line like `iniPrm := []float32{10, 9}` which would set 10 and 9 as the initial values. `maxsteps` is the maximum number of steps. You then run the simplex with a line like

 result, err := s.Run(2)

which would run the simplex two times with `maxsteps` steps each time. To get to the optimized parameters, you would `fmt.Println (result.BestPrm)`. To check if convergence was reached, you might say

 if result.StopReason != simplex.Converged {
     fmt.Println("Did not converge")
 }

Really, one shoul just `go doc -all` to see many more options.