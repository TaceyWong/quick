package starter

// StarterOption options for create Starter
type StarterOption struct {
	Template          string
	Checkout          string
	NoInput           bool
	ExtraContext      interface{}
	Replay            string
	OverwriteIfExists bool
	OutputDir         string
	ConfigFile        string
	UseDefaultConfig  bool
	Password          string
	Directory         string
	SkipIfFileExists  bool
	AcceptHooks       bool
}

// Starter main struct&executor
type Starter struct {
}

// NewStarter create a instance of Starter
func NewStarter(opt *StarterOption) *Starter {
	return &Starter{}
}

// List  List installed (locally cloned) templates.
func (s *Starter) List() error {
	return nil
}

// Touch just like linux-touch(re-clone template)
func (s *Starter) Touch() error {
	return nil
}

// Run run full flow
func (s *Starter) Run() error {
	// Step-1: rake config
	// Step-2: clone|use-local template
	// Step-3: ask-flow
	// Step-4: generate project

	return nil
}

// Clone clone|use-local template
func (s *Starter) Clone() error {
	return nil
}

// Ask ask-flow
func (s *Starter) Ask() error {
	return nil
}

// Generate generate the project
func (s *Starter) Generate() error {
	return nil
}

func doPreHook() error {
	return nil
}

func doAfterHook() error {
	return nil
}
