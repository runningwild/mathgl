include $(GOROOT)/src/Make.inc

TARG=github.com/arbaal/mathgl

OFILES_arm=\
	fsqrt32_arm.$O\

OFILES_amd64=\
	fsqrt32_amd64.$O\

OFILES_386=\
	fsqrt32_386.$O\

OFILES=\
	$(OFILES_$(GOARCH))

ALLGOFILES=\
	fsincos32.go\
	mat3.go\
	func.go\
	vec2.go\

NOGOFILES=\
	$(subst _$(GOARCH).$O,.go,$(OFILES_$(GOARCH)))

GOFILES=\
	$(filter-out $(NOGOFILES),$(ALLGOFILES))\
	$(subst .go,_decl.go,$(NOGOFILES))\

include $(GOROOT)/src/Make.pkg
