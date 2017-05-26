// This file was generated by https://github.com/ChimeraCoder/gojson.
package main

type ebConfig struct {
	Commands struct {
		CMD_AppDeploy struct {
			RefreshManifest string `json:"refresh_manifest"`
			Stages          []struct {
				Actions []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"actions"`
				Name string `json:"name"`
			} `json:"stages"`
		} `json:"CMD-AppDeploy"`
		CMD_BundleLogs struct {
			PersistentConfiguration string `json:"persistent_configuration"`
			Stages                  []struct {
				Actions []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"actions"`
				Name string `json:"name"`
			} `json:"stages"`
		} `json:"CMD-BundleLogs"`
		CMD_ConfigDeploy struct {
			RefreshManifest string `json:"refresh_manifest"`
			Stages          []struct {
				Actions []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"actions"`
				Name string `json:"name"`
			} `json:"stages"`
		} `json:"CMD-ConfigDeploy"`
		CMD_ImmutableDeploymentFlip struct {
			PersistentConfiguration string `json:"persistent_configuration"`
			Stages                  []struct {
				Actions []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"actions"`
				Name string `json:"name"`
			} `json:"stages"`
		} `json:"CMD-ImmutableDeploymentFlip"`
		CMD_PreInit struct {
			RefreshManifest string `json:"refresh_manifest"`
			Stages          []struct {
				Actions []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"actions"`
				Name string `json:"name"`
			} `json:"stages"`
		} `json:"CMD-PreInit"`
		CMD_PublishLogs struct {
			PersistentConfiguration string `json:"persistent_configuration"`
			Stages                  []struct {
				Actions []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"actions"`
				Name string `json:"name"`
			} `json:"stages"`
		} `json:"CMD-PublishLogs"`
		CMD_RestartAppServer struct {
			Stages []struct {
				Actions []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"actions"`
				Name string `json:"name"`
			} `json:"stages"`
		} `json:"CMD-RestartAppServer"`
		CMD_SelfStartup struct {
			Stages []struct {
				Actions []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"actions"`
				Name string `json:"name"`
			} `json:"stages"`
		} `json:"CMD-SelfStartup"`
		CMD_Startup struct {
			Stages []struct {
				Actions []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"actions"`
				Name string `json:"name"`
			} `json:"stages"`
		} `json:"CMD-Startup"`
		CMD_SystemTailLogs struct {
			PersistentConfiguration string `json:"persistent_configuration"`
			Stages                  []struct {
				Actions []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"actions"`
				Name string `json:"name"`
			} `json:"stages"`
		} `json:"CMD-SystemTailLogs"`
		CMD_TailLogs struct {
			PersistentConfiguration string `json:"persistent_configuration"`
			Stages                  []struct {
				Actions []struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"actions"`
				Name string `json:"name"`
			} `json:"stages"`
		} `json:"CMD-TailLogs"`
	} `json:"commands"`
	Container struct {
		AppAssetDir        string   `json:"app_asset_dir"`
		AppDeployDir       string   `json:"app_deploy_dir"`
		AppLogDir          string   `json:"app_log_dir"`
		AppPidDir          string   `json:"app_pid_dir"`
		AppStagingDir      string   `json:"app_staging_dir"`
		AppUser            string   `json:"app_user"`
		CommonLogList      []string `json:"common_log_list"`
		ConfigStagingDir   string   `json:"config_staging_dir"`
		DefaultLogList     []string `json:"default_log_list"`
		GemDir             string   `json:"gem_dir"`
		HTTPPort           string   `json:"http_port"`
		InstancePort       string   `json:"instance_port"`
		LogGroupNamePrefix string   `json:"log_group_name_prefix"`
		NodeInstallDir     string   `json:"node_install_dir"`
		NodeVersion        string   `json:"node_version"`
		PumaLogDir         string   `json:"puma_log_dir"`
		PumaPidDir         string   `json:"puma_pid_dir"`
		PumaVersion        string   `json:"puma_version"`
		RubyVersion        string   `json:"ruby_version"`
		ScriptDir          string   `json:"script_dir"`
		SourceBundle       string   `json:"source_bundle"`
		SupportDir         string   `json:"support_dir"`
		TarballURL         string   `json:"tarball_url"`
	} `json:"container"`
	Optionsettings struct {
		Environment                      []string `json:"aws:elasticbeanstalk:application:environment"`
		Aws_elasticbeanstalk_hostmanager struct {
			LogPublicationControl string `json:"LogPublicationControl"`
		} `json:"aws:elasticbeanstalk:hostmanager"`
	} `json:"optionsettings"`
	System struct {
		AWSEBAgentID          string `json:"AWSEBAgentId"`
		AWSEBReferrerID       string `json:"AWSEBReferrerId"`
		LogPublicationControl string `json:"LogPublicationControl"`
	} `json:"system"`
}
