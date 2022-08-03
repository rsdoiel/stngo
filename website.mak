#
# Makefile for running pandoc on all Markdown docs ending in .md
#
PROJECT = stn

MD_PAGES = $(shell ls -1 *.md | grep -v "nav.md")

HTML_PAGES = $(shell ls -1 *.md | grep -v "nav.md" | sed -E 's/.md/.html/g')

build: $(HTML_PAGES) $(MD_PAGES) license.html stn.html stnfilter.html stnparse.html stnreport.html

$(HTML_PAGES): $(MD_PAGES) .FORCE
	pandoc --metadata title=$(basename $@) -s --to html5 $(basename $@).md -o $(basename $@).html \
	    --template=page.tmpl
	@if [ $@ = "README.html" ]; then mv README.html index.html; fi

license.html: LICENSE
	pandoc --metadata title="$(PROJECT): License" -s --from Markdown --to html5 LICENSE -o license.html \
	    --template=page.tmpl

stn.html: docs/stn.md
	pandoc --metadata title="$(PROJECT): simple timesheet notation" -s --from Markdown --to html5 docs/stn.md -o stn.html --template page.tmpl
	
stnfilter.html: docs/stnfilter.md
	pandoc --metadata title="$(PROJECT): stnfilter" -s --from Markdown --to html5 docs/stnfilter.md -o stnfilter.html --template page.tmpl

stnparse.html: docs/stnparse.md
	pandoc --metadata title="$(PROJECT): stnparse" -s --from Markdown --to html5 docs/stnparse.md -o stnparse.html --template page.tmpl

stnreport.html: docs/stnreport.md
	pandoc --metadata title="$(PROJECT): stnreport" -s --from Markdown --to html5 docs/stnreport.md -o stnreport.html --template page.tmpl

clean:
	@if [ -f index.html ]; then rm *.html; fi

.FORCE:
