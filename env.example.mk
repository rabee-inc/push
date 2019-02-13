LOCAL_PROJECT_ID = 'develop-xxxx-rabee-jp'
STAGING_PROJECT_ID = 'staging2-salontia-rabee-jp'
PRODUCTION_PROJECT_ID = 'salontia-rabee-jp'

define apps
	$(call init,local,push)
	$(call init,staging,push)
	$(call init,production,push)
endef
