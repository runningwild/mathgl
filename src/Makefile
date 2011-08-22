include $(GOROOT)/src/Make.inc

TARG=mathgl

OFILES_amd64=\
	fsqrt32_amd64.$O\

OFILES=\
	$(OFILES_$(GOARCH))

ALLGOFILES=\
	fsqrt32.go\
	fsqrt32_port.go\
	mat3.go\
	util.go\
	vec2.go\

NOGOFILES=\
	$(subst _$(GOARCH).$O,.go,$(OFILES_$(GOARCH)))

GOFILES=\
	$(filter-out $(NOGOFILES),$(ALLGOFILES))\
	$(subst .go,_decl.go,$(NOGOFILES))\

include $(GOROOT)/src/Make.pkg
