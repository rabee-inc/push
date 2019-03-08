GOPHER = 'ʕ◔ϖ◔ʔ'

.PHONY: hello init run deploy

hello:
	@echo Hello go project ${GOPHER}

# 準備
init:
	${call init}

# [GAE] アプリの実行
run:
	${call init}
	${call run,local,${app}}

run-staging:
	${call init}
	${call run,staging,${app}}

run-production:
	${call init}
	${call run,production,${app}}

# [GAE] アプリのデプロイ
deploy:
	${call init}
	${call deploy,staging,${app},${STAGING_PROJECT_ID}}

deploy-production:
	${call init}
	${call deploy,production,${app},${PRODUCTION_PROJECT_ID}}

# [GAE] ディスパッチ設定をデプロイ
deploy-dispatch:
	${call deploy-config,staging,dispatch.yaml,${STAGING_PROJECT_ID}}

deploy-dispatch-production:
	${call deploy-config,production,dispatch.yaml,${PRODUCTION_PROJECT_ID}}

# [GAE] Cron設定をデプロイ
deploy-cron:
	${call deploy-config,staging,cron.yaml,${STAGING_PROJECT_ID}}

deploy-cron-production:
	${call deploy-config,production,cron.yaml,${PRODUCTION_PROJECT_ID}}

# [GAE] Queue設定をデプロイ
deploy-queue:
	${call deploy-config,staging,queue.yaml,${STAGING_PROJECT_ID}}

deploy-queue-production:
	${call deploy-config,production,queue.yaml,${PRODUCTION_PROJECT_ID}}

# [GAE] Datastoreの複合インデックス定義をデプロイ
deploy-index:
	${call deploy-config,staging,index.yaml,${STAGING_PROJECT_ID}}

deploy-index-production:
	${call deploy-config,production,index.yaml,${PRODUCTION_PROJECT_ID}}

# [Firestore] 全データ削除
firestore-delete:
	${call firestore-delete,${LOCAL_PROJECT_ID}}

firestore-delete-staging:
	${call firestore-delete,${STAGING_PROJECT_ID}}

# マクロ
define init
	@go run ./command/init/main.go 
endef

define run
	@go run ./command/run/main.go -env $1 -app $2
endef

define deploy
	@gcloud app deploy -q deploy/appengine/$1/$2/app.yaml --project $3
endef

define deploy-config
	@gcloud app deploy -q deploy/appengine/$1/push/$2 --project $3
endef

define firestore-delete
	firebase firestore:delete --all-collections --project $1
endef

include deploy/env.mk