# Run your Go OpenGL app with NVIDIA card

This package exports the `NvOptimusEnablement` variable. When NVIDIA Optimus drivers 
detects that variable (set to 1), it runs your app on NVIDIA card instead of the 
integrated one. This solution was found by 
[this person](https://groups.google.com/d/msg/golang-nuts/7OHZcXUegF0/dn8Ni__KAAAJ)

This package will have no effect on non windows systems

# Usage / Installation

All you have to do is to `go get` this package:

```
go get github.com/aqatl/nvidiaoptimus
```

And then just import it:

```
import (
	_ "github.com/aqatl/nvidiaoptimus"
)
```

That's it!

