%.zip:
	cd $(basename $@) && \
		pipenv requirements > requirements.txt && \
		pip install --target package -r requirements.txt && \
		cd package && \
		zip -r ../../$@ . && \
		cd .. && \
		zip -g ../$@ lambda_function.py && \
		rm -rf package requirements.txt
