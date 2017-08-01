package main

import (
	"encoding/xml"
	"os"

	"github.com/bndr/gojenkins"

	"fmt"
)

var (
	jenkins *gojenkins.Jenkins
	s       []InnerJobs
	m       InnerJob
)

type InnerJobs struct {
	Raw      []InnerJob
	NodeName string
	IsFolder bool
}
type InnerJob struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Color string `json:"color"`
}

//
func main() {
	// cliTest()
	// beego.Run()
	TestInit()
	// TestGetAllJobs()
	// TestCreateJobs()
}

func TestInit() {
	jenkins = gojenkins.CreateJenkins("http://localhost:8080", "admin", "admin")
	_, err := jenkins.Init()
	if err == nil {
		fmt.Println("Jenkins Initialization success")
	} else {
		fmt.Errorf("Jenkins Initialization fail", err)
	}
	jobs, err := jenkins.GetAllJobs()
	if err != nil {
		fmt.Println("GetAllJobNames went wrong")
		return
	}
	fmt.Println(jobs)
	job := jobs[2]
	build, _ := job.GetBuild(15)

	context := build.GetConsoleOutput()
	// context := build.GetConsoleOutputByNum()
	fmt.Println(context)
	// code, _ := jenkins.Poll()
	// fmt.Println(code)
	// jenkins, _ := jenkins.Info()
	// fmt.Println(jenkins)
}

func TestGetAllJobs() {
	jobs, err := jenkins.GetAllJobNames()
	if err != nil {
		fmt.Println("GetAllJobNames went wrong")
		return
	}
	getFolderJobs(jobs, "", "")
	fmt.Println(s)
}

func getFolderJobs(jobs []gojenkins.InnerJob, prefix string, nodeName string) {
	var prefixTemp string
	var arr []InnerJob
	isFolder := false
	for _, job := range jobs {
		if job.Color == "" {
			isFolder = true
			if prefix != "" {
				prefixTemp = prefix + "/job/"
			}
			folder, err := jenkins.GetFolder(prefixTemp + job.Name)
			if err != nil {
				fmt.Println("GetFolder went wrong")
			}
			jobs := folder.Raw.Jobs
			if len(jobs) != 0 {
				getFolderJobs(jobs, prefixTemp+job.Name, job.Name)
			}
		} else {
			cr := InnerJob{Name: job.Name, Color: job.Color, Url: job.Url}
			arr = append(arr, cr)
		}
	}
	s = append(s, InnerJobs{Raw: arr, IsFolder: isFolder, NodeName: nodeName})
}

func TestCreateJobs() {
	jobID := "testXML04"
	job_data := getFileAsString("jobTest.xml")
	job, err := jenkins.CreateJob(job_data, jobID)
	if err != nil {
		fmt.Errorf("CreateJob fail", err)
	} else {
		fmt.Println("CreateJob success")
	}
	fmt.Println(jobID, job.GetName())
}

func getFileAsString(path string) string {

	// buf, err := ioutil.ReadFile("/Users/admin/Documents/go/src/test/" + path)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("getFileAsString success")
	// }
	// Header := `<?xml version="1.0" encoding="UTF-8"?>`
	return xml.Header + string(createXml())
}

type FlowDefinition struct {
	XMLName          xml.Name   `xml:"flow-definition"`
	Plugin           string     `xml:"plugin,attr"`
	Description      string     `xml:"description"`
	KeepDependencies bool       `xml:"keepDependencies"`
	Properties       Properties `xml:"properties"`
	Definition       Definition `xml:"definition"`
	Triggers         string     `xml:"triggers"`
	Disabled         bool       `xml:"disabled"`
}
type Properties struct {
	XMLName                      xml.Name                     `xml:"properties"`
	GithubProjectProperty        GithubProjectProperty        `xml:"com.coravy.hudson.plugins.github.GithubProjectProperty"`
	ParametersDefinitionProperty ParametersDefinitionProperty `xml:"hudson.model.ParametersDefinitionProperty"`
	PipelineTriggersJobProperty  PipelineTriggersJobProperty  `xml:"org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty"`
}

type GithubProjectProperty struct {
	XMLName     xml.Name `xml:"com.coravy.hudson.plugins.github.GithubProjectProperty"`
	Plugin      string   `xml:"plugin,attr"`
	ProjectUrl  string   `xml:"projectUrl"`
	DisplayName string   `xml:"displayName"`
}
type ParametersDefinitionProperty struct {
	XMLName              xml.Name             `xml:"hudson.model.ParametersDefinitionProperty"`
	ParameterDefinitions ParameterDefinitions `xml:"parameterDefinitions"`
}
type ParameterDefinitions struct {
	XMLName                   xml.Name                    `xml:"parameterDefinitions"`
	StringParameterDefinition []StringParameterDefinition `xml:"hudson.model.StringParameterDefinition"`
}
type StringParameterDefinition struct {
	XMLName      xml.Name `xml:"hudson.model.StringParameterDefinition"`
	Name         string   `xml:"name"`
	Description  string   `xml:"description"`
	DefaultValue string   `xml:"defaultValue"`
}
type PipelineTriggersJobProperty struct {
	XMLName  xml.Name `xml:"org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty"`
	Triggers Triggers `xml:"triggers"`
}
type Triggers struct {
	XMLName      xml.Name     `xml:"triggers"`
	GhprbTrigger GhprbTrigger `xml:"org.jenkinsci.plugins.ghprb.GhprbTrigger"`
}
type GhprbTrigger struct {
	XMLName                              xml.Name                `xml:"org.jenkinsci.plugins.ghprb.GhprbTrigger"`
	Plugin                               string                  `xml:"plugin,attr"`
	Spec                                 string                  `xml:"spec"`
	ConfigVersion                        string                  `xml:"configVersion"`
	Adminlist                            string                  `xml:"adminlist"`
	AllowMembersOfWhitelistedOrgsAsAdmin bool                    `xml:"allowMembersOfWhitelistedOrgsAsAdmin"`
	Orgslist                             string                  `xml:"orgslist"`
	Cron                                 string                  `xml:"cron"`
	BuildDescTemplate                    string                  `xml:"buildDescTemplate"`
	OnlyTriggerPhrase                    bool                    `xml:"onlyTriggerPhrase"`
	UseGitHubHooks                       bool                    `xml:"useGitHubHooks"`
	PermitAll                            bool                    `xml:"permitAll"`
	Whitelist                            string                  `xml:"whitelist"`
	AutoCloseFailedPullRequests          bool                    `xml:"autoCloseFailedPullRequests"`
	DisplayBuildErrorsOnDownstreamBuilds bool                    `xml:"displayBuildErrorsOnDownstreamBuilds"`
	WhiteListTargetBranches              WhiteListTargetBranches `xml:"whiteListTargetBranches"`
	BlackListTargetBranches              BlackListTargetBranches `xml:"blackListTargetBranches"`
	GitHubAuthId                         string                  `xml:"gitHubAuthId"`
	TriggerPhrase                        string                  `xml:"triggerPhrase"`
	SkipBuildPhrase                      string                  `xml:"skipBuildPhrase"`
	BlackListCommitAuthor                string                  `xml:"blackListCommitAuthor"`
	BlackListLabels                      string                  `xml:"blackListLabels"`
	WhiteListLabels                      string                  `xml:"whiteListLabels"`
	IncludedRegions                      string                  `xml:"includedRegions"`
	ExcludedRegions                      string                  `xml:"excludedRegions"`
	Extensions                           Extensions              `xml:"extensions"`
}
type BlackListTargetBranches struct {
	XMLName     xml.Name    `xml:"blackListTargetBranches"`
	GhprbBranch GhprbBranch `xml:"org.jenkinsci.plugins.ghprb.GhprbBranch"`
}
type WhiteListTargetBranches struct {
	XMLName     xml.Name    `xml:"whiteListTargetBranches"`
	GhprbBranch GhprbBranch `xml:"org.jenkinsci.plugins.ghprb.GhprbBranch"`
}
type GhprbBranch struct {
	XMLName xml.Name `xml:"org.jenkinsci.plugins.ghprb.GhprbBranch"`
	Branche string   `xml:"branch"`
}

type Extensions struct {
	XMLName           xml.Name          `xml:"extensions"`
	GhprbSimpleStatus GhprbSimpleStatus `xml:"org.jenkinsci.plugins.ghprb.extensions.status.GhprbSimpleStatus"`
}
type GhprbSimpleStatus struct {
	XMLName             xml.Name `xml:"org.jenkinsci.plugins.ghprb.extensions.status.GhprbSimpleStatus"`
	CommitStatusContext string   `xml:"commitStatusContext"`
	TriggeredStatus     string   `xml:"triggeredStatus"`
	StartedStatus       string   `xml:"startedStatus"`
	StatusUrl           string   `xml:"statusUrl"`
	AddTestResults      bool     `xml:"addTestResults"`
}
type Definition struct {
	XMLName     xml.Name `xml:"definition"`
	Class       string   `xml:"class,attr"`
	Plugin      string   `xml:"plugin,attr"`
	Scm         Scm      `xml:"scm"`
	ScriptPath  string   `xml:"scriptPath"`
	Lightweight bool     `xml:"lightweight"`
}
type Scm struct {
	XMLName                           xml.Name          `xml:"scm"`
	Class                             string            `xml:"class,attr"`
	Plugin                            string            `xml:"plugin,attr"`
	ConfigVersion                     string            `xml:"configVersion"`
	UserRemoteConfigs                 UserRemoteConfigs `xml:"userRemoteConfigs"`
	Branches                          Branches          `xml:"branches"`
	DoGenerateSubmoduleConfigurations bool              `xml:"doGenerateSubmoduleConfigurations"`
	SubmoduleCfg                      SubmoduleCfg      `xml:"submoduleCfg"`
	Extensions                        string            `xml:"extensions"`
}
type SubmoduleCfg struct {
	XMLName xml.Name `xml:"submoduleCfg"`
	Class   string   `xml:"class,attr"`
}
type UserRemoteConfigs struct {
	XMLName          xml.Name         `xml:"userRemoteConfigs"`
	UserRemoteConfig UserRemoteConfig `xml:"hudson.plugins.git.UserRemoteConfig"`
}
type UserRemoteConfig struct {
	XMLName       xml.Name `xml:"hudson.plugins.git.UserRemoteConfig"`
	Url           string   `xml:"url"`
	CredentialsId string   `xml:"credentialsId"`
}
type Branches struct {
	XMLName    xml.Name   `xml:"branches"`
	BranchSpec BranchSpec `xml:"hudson.plugins.git.BranchSpec"`
}
type BranchSpec struct {
	XMLName xml.Name `xml:"hudson.plugins.git.BranchSpec"`
	Name    string   `xml:"name"`
}

func createXml() []byte {
	f := &FlowDefinition{
		Plugin:           "workflow-job@2.12.1",
		Description:      "",
		KeepDependencies: false,
		Disabled:         false}
	// f.Properties = Properties{}
	f.Properties.GithubProjectProperty = GithubProjectProperty{Plugin: "github@1.27.0", ProjectUrl: "https://github.com/baomengjiang/test/"}
	// f.Properties.ParametersDefinitionProperty = ParametersDefinitionProperty{}
	f.Properties.ParametersDefinitionProperty.ParameterDefinitions.StringParameterDefinition = append(f.Properties.ParametersDefinitionProperty.ParameterDefinitions.StringParameterDefinition, StringParameterDefinition{Name: "sha1", Description: "", DefaultValue: "*/master"})
	// f.Properties.PipelineTriggersJobProperty = PipelineTriggersJobProperty{}
	// f.Properties.PipelineTriggersJobProperty.Triggers = Triggers{}
	f.Properties.PipelineTriggersJobProperty.Triggers.GhprbTrigger = GhprbTrigger{Plugin: "ghprb@1.39.0", Spec: "H/5 * * * *", ConfigVersion: "3", Adminlist: "baomengjiang", AllowMembersOfWhitelistedOrgsAsAdmin: false, Cron: "H/5 * * * *", OnlyTriggerPhrase: false, UseGitHubHooks: true, PermitAll: false, AutoCloseFailedPullRequests: false, DisplayBuildErrorsOnDownstreamBuilds: false, GitHubAuthId: "ccadcc64-5d20-4bcc-8358-a22adbd157ee", SkipBuildPhrase: `.*\[skip\W+ci\].*`}
	// f.Properties.PipelineTriggersJobProperty.Triggers.GhprbTrigger.Extensions = Extensions{}
	f.Properties.PipelineTriggersJobProperty.Triggers.GhprbTrigger.Extensions.GhprbSimpleStatus = GhprbSimpleStatus{AddTestResults: false}
	f.Definition = Definition{Class: "org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition", Plugin: "workflow-cps@2.36.1", ScriptPath: "Jenkinsfile", Lightweight: false}
	f.Definition.Scm = Scm{Class: "hudson.plugins.git.GitSCM", Plugin: "git@3.4.0", ConfigVersion: "2", DoGenerateSubmoduleConfigurations: false}
	f.Definition.Scm.SubmoduleCfg = SubmoduleCfg{Class: "list"}
	// f.Definition.Scm.UserRemoteConfigs = UserRemoteConfigs{}
	f.Definition.Scm.UserRemoteConfigs.UserRemoteConfig = UserRemoteConfig{Url: "https://github.com/baomengjiang/test", CredentialsId: "462799f9-a326-40e3-a425-026c320fc3be"}
	// f.Definition.Scm.Branches = Branches{}
	f.Definition.Scm.Branches.BranchSpec = BranchSpec{Name: "*/master"}
	output, err := xml.MarshalIndent(f, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
	return output
}
