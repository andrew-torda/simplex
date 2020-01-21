// License ? what license ?

/*
Package simplex implements the Nelder-Mead simplex optimizer.
It follows the original paper
   J.A. Nelder, R. Mead (1965), Comp. J., 7, 308-313

but some of the nomenclature comes from
   J.C. Lagarias, J.A. Reeds, M.H. Wright, and P.E. Wright (1998),
   SIAM J. Optim, 9, 112-147.

and a tiny bit from 
 Press, W.H., Teukolsky, S.A., Vetterling, W.T., Flannery, B.P.,
 Numerical Recipes in C., Cambridge University Press, 1992

The name of the amotry() function comes from numerical recipes,
but the formulae for moving the highest point around are taken from
the primary references.

It has a couple of frills.
 1. There are two ways to initialise. You can use the classic version as in
    numerical recipes. Works well on my very artificial examples which have
    some funny symmetries. This is the default.
    Alternatively, you can the initial points in each dimension spread evenly
    over the allowed values and surrounding the initial parameter values.
    To do this, call

   s.IniType (simplex.IniPntSpread)

    When `s` has the type splxctrl.
    

 2. It allows a vector of minimum and maximum values. It will reject
    moves if they go beyond these boundaries. This is done by wrapping
    the cost function. If you do not exceed the bounds, the wrapper just
    returns the original cost function. If a point exceeds a bound in any
    dimension, the cost function is set to NaN, so the point will be
    rejected.

Could be improved

 * we call a full sort after each cycle. This is not necessary. One
   only needs a list with the highest, next-highest and best points.
 * We re-calculate the centroid on each cycle. This is necessary, but
   the version in numerical recipes does it by removing the old best
   value and adding in the new one. In a few dimensions, it makes no
   difference. In many dimensions, one could argue the method in numerical
   recipes will be faster.
*/
package simplex
