OBJECTS = $(foreach item,$(shell find . -maxdepth 1 -type d -not -path . -not -name .pytest_cache),$(item).zip)

all: $(OBJECTS)

%.zip:
	cd $(basename $@) && \
		pipenv requirements > requirements.txt && \
		pip install --target package -r requirements.txt && \
		cd package && \
		zip -r ../../$@ . && \
		cd .. && \
		zip -g ../$@ lambda_function.py && \
		rm -rf package requirements.txt

clean:
	rm -f *.zip
	find . -name package -type d -exec rm -rf {} +
	find . -name requirements.txt -type f -exec rm -rf {} \;
