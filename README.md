# simplex
Nelder Meade simplex optimization method.

The basic interface is simple. You ask for a new object and then say Run. There is an absurd number of options. You can specify upper and lower bounds for parameters individually, convergence requirements for each parameter and there are two different methods for setting up the initial vertices. The most complicated option invokes a second cost function on the data each time a new best point is found. If one is doing an optimisation for some kind of fitting, you can have a standard cost function that acts on most of the data, but a second function that acts on part of the data for testing (to see recall versus generalisation or overfitting).
