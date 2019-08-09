package classpath

import "os"
import "path/filepath"

type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}
