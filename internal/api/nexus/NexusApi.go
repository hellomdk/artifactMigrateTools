package nexus

import (
	"artifactMigrateTools/internal/config"
	"artifactMigrateTools/internal/util"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/**
探活
client.Ping()
*/
func (n *Nexus) Ping(context *config.Context) (string, error) {
	url := "/service/metrics/ping"
	statusCode, body, _ := util.Client.GetString(n.HttpClient.BaseURL+url, n.HttpClient.Header, nil)
	if statusCode != 200 {
		context.Loggers.SendLoggerError(fmt.Sprint("API测试连接失败: ", statusCode), nil)
	}
	return body, nil
}

/**
获取bucket 文件列表
client.CreateArtifact("Gradle_Test")
*/
func (n *Nexus) GetItemList(context *config.Context, repoKey string) NexusFileItem {
	url := "/service/rest/v1/components1?repository={repoKey}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)

	var cacheConfig = &NexusFileItem{}
	statusCode, _ := util.Client.Get(n.HttpClient.BaseURL+url, n.HttpClient.Header, nil, cacheConfig)

	if statusCode != 200 {
		context.Loggers.SendLoggerError(fmt.Sprintf("API获取bucket文件列表【%s】失败: %d", url, statusCode), nil)
	}

	return *cacheConfig
}

/**
读取manifests
client.CreateArtifact("Gradle_Test")
*/
func (n *Nexus) ReadManifests(context *config.Context, repoKey, repoPath string) NexusDockerManifests {
	url := "/repository/{repoKey}/{repoPath}"
	url = strings.Replace(url, "{repoKey}", repoKey, 1)
	url = strings.Replace(url, "{repoPath}", repoPath, 1)

	var cacheConfig = &NexusDockerManifests{}
	_, err := util.Client.Get(n.HttpClient.BaseURL+url, n.HttpClient.Header, nil, cacheConfig)

	if err != nil {
		context.Loggers.SendLoggerError("API获取manifests失败: ", err)
	}

	return *cacheConfig
}

func (n *Nexus) RegisterScript(context *config.Context) bool {
	url := "/service/rest/v1/script"
	body := "{\"name\":\"artifactorymigrator\",\"type\":\"groovy\",\"content\":\"import groovy.json.JsonBuilder\\nimport org.sonatype.nexus.security.user.UserSearchCriteria\\nimport org.apache.shiro.authz.permission.*\\n\\ndef getRepoData() {\\n    def repoman = container.lookup('org.sonatype.nexus.repository.manager.RepositoryManager')\\n    def repolist = []\\n    for (repo in repoman.browse()) {\\n        def repoitem = [:], configitem = [:]\\n        repoitem.type = repo.type.value\\n        repoitem.format = repo.format.value\\n        repoitem.name = repo.name\\n        repoitem.url = repo.url\\n        def config = repo.configuration\\n        configitem.name = config.repositoryName\\n        configitem.recipe = config.recipeName\\n        configitem.online = config.isOnline()\\n        configitem.attributes = config.attributes\\n        repoitem.config = configitem\\n        repolist << repoitem\\n    }\\n    return repolist\\n}\\n\\ndef getBlobstoreData() {\\n    def blobman = container.lookup('org.sonatype.nexus.blobstore.api.BlobStoreManager')\\n    def blobstorelist = []\\n    for (blobstore in blobman.browse()) {\\n        def blobstoreitem = [:]\\n        def config = blobstore.blobStoreConfiguration\\n        blobstoreitem.name = config.name\\n        blobstoreitem.type = config.type\\n        blobstoreitem.attributes = config.attributes\\n        blobstorelist << blobstoreitem\\n    }\\n    return blobstorelist\\n}\\n\\ndef getUserData() {\\n    def secsys = container.lookup('org.sonatype.nexus.security.SecuritySystem')\\n    def userlist = []\\n    for (user in secsys.searchUsers(new UserSearchCriteria())) {\\n        def useritem = [:], rolesitems = []\\n        useritem.id = user.userId\\n        useritem.firstname = user.firstName\\n        useritem.lastname = user.lastName\\n        useritem.email = user.emailAddress\\n        useritem.source = user.source\\n        for (role in user.roles) {\\n            def roleitem = [:]\\n            roleitem.id = role.roleId\\n            roleitem.source = role.source\\n            rolesitems << roleitem\\n        }\\n        useritem.roles = rolesitems\\n        useritem.status = user.status.name()\\n        useritem.readonly = user.isReadOnly()\\n        useritem.version = user.version\\n        userlist << useritem\\n    }\\n    return userlist\\n}\\n\\ndef getGroupData() {\\n    def secsys = container.lookup('org.sonatype.nexus.security.SecuritySystem')\\n    def grouplist = []\\n    for (group in secsys.listRoles()) {\\n        def groupitem = [:]\\n        groupitem.id = group.roleId\\n        groupitem.name = group.name\\n        groupitem.source = group.source\\n        groupitem.roles = group.roles\\n        groupitem.privileges = group.privileges\\n        groupitem.description = group.description\\n        groupitem.readonly = group.readOnly\\n        groupitem.version = group.version\\n        grouplist << groupitem\\n    }\\n    return grouplist\\n}\\n\\ndef getPermissionData() {\\n    def secsys = container.lookup('org.sonatype.nexus.security.SecuritySystem')\\n    def permlist = []\\n    for (perm in secsys.listPrivileges()) {\\n        def permitem = [:]\\n        permitem.id = perm.id\\n        permitem.name = perm.name\\n        permitem.description = perm.description\\n        permitem.type = perm.type\\n        permitem.properties = perm.properties\\n        permitem.readonly = perm.isReadOnly()\\n        permitem.version = perm.version\\n        permitem.perm = perm.permission.parts\\n        permlist << permitem\\n    }\\n    return permlist\\n}\\n\\ndef getSelectorData() {\\n    def selman = container.lookup('org.sonatype.nexus.selector.SelectorManager')\\n    def sellist = []\\n    if (selman == null) return sellist\\n    for (sel in selman.browse()) {\\n        def selitem = [:]\\n        selitem.name = sel.name\\n        selitem.type = sel.type\\n        selitem.description = sel.description\\n        selitem.attributes = sel.attributes\\n        sellist << selitem\\n    }\\n    return sellist\\n}\\n\\ndef getLdapData() {\\n    def confman = container.lookup('org.sonatype.nexus.ldap.persist.LdapConfigurationManager')\\n    def conflist = []\\n    for (conf in confman.listLdapServerConfigurations()) {\\n        def confitem = [:]\\n        confitem.id = conf.id\\n        confitem.name = conf.name\\n        confitem.order = conf.order\\n        confitem.searchBase = conf.connection.searchBase\\n        confitem.systemUsername = conf.connection.systemUsername\\n        confitem.systemPassword = conf.connection.systemPassword\\n        confitem.authScheme = conf.connection.authScheme\\n        confitem.useTrustStore = conf.connection.useTrustStore\\n        confitem.saslRealm = conf.connection.saslRealm\\n        confitem.connectionTimeout = conf.connection.connectionTimeout\\n        confitem.connectionRetryDelay = conf.connection.connectionRetryDelay\\n        confitem.maxIncidentsCount = conf.connection.maxIncidentsCount\\n        confitem.protocol = conf.connection.host.protocol.name()\\n        confitem.hostName = conf.connection.host.hostName\\n        confitem.port = conf.connection.host.port\\n        confitem.emailAddressAttribute = conf.mapping.emailAddressAttribute\\n        confitem.ldapGroupsAsRoles = conf.mapping.ldapGroupsAsRoles\\n        confitem.groupBaseDn = conf.mapping.groupBaseDn\\n        confitem.groupIdAttribute = conf.mapping.groupIdAttribute\\n        confitem.groupMemberAttribute = conf.mapping.groupMemberAttribute\\n        confitem.groupMemberFormat = conf.mapping.groupMemberFormat\\n        confitem.groupObjectClass = conf.mapping.groupObjectClass\\n        confitem.userPasswordAttribute = conf.mapping.userPasswordAttribute\\n        confitem.userIdAttribute = conf.mapping.userIdAttribute\\n        confitem.userObjectClass = conf.mapping.userObjectClass\\n        confitem.ldapFilter = conf.mapping.ldapFilter\\n        confitem.userBaseDn = conf.mapping.userBaseDn\\n        confitem.userRealNameAttribute = conf.mapping.userRealNameAttribute\\n        confitem.userSubtree = conf.mapping.userSubtree\\n        confitem.groupSubtree = conf.mapping.groupSubtree\\n        confitem.userMemberOfAttribute = conf.mapping.userMemberOfAttribute\\n        conflist << confitem\\n    }\\n    return conflist\\n}\\n\\ndef getData() {\\n    def data = [:]\\n    if (true) {\\n        data.users = []\\n        data.groups = []\\n        data.privs = []\\n        data.selectors = []\\n        data.ldaps = []\\n    } else {\\n        data.users = userData\\n        data.groups = groupData\\n        data.privs = permissionData\\n        data.selectors = selectorData\\n        data.ldaps = ldapData\\n    }\\n    data.repos = repoData\\n    data.blobstores = blobstoreData\\n    return new JsonBuilder(data).toPrettyString()\\n}\\n\\nreturn data\\n\"}"
	var cacheConfig = &NexusDataDto{}
	statusCode, _ := util.Client.Post(n.HttpClient.BaseURL+url, n.HttpClient.Header, body, cacheConfig)

	if statusCode != 200 {
		context.Loggers.SendLoggerError(fmt.Sprint("API注册nexus script失败: ", statusCode), nil)
		return false
	}

	return true
}

func (n *Nexus) GetNexusData(context *config.Context, scriptName string) NexusData {
	url := "/service/rest/v1/script/{scriptName}/run"
	url = strings.Replace(url, "{scriptName}", scriptName, 1)

	var cacheConfig = &NexusDataDto{}
	statusCode, _ := util.Client.Post(n.HttpClient.BaseURL+url, n.HttpClient.Header, nil, cacheConfig)

	if statusCode != 200 {
		context.Loggers.SendLoggerError(fmt.Sprint("API获取nexus数据失败: ", statusCode), nil)
	}
	jsonStr := cacheConfig.Result
	nexusData := NexusData{}
	err := json.Unmarshal([]byte(jsonStr), &nexusData)
	if err != nil {
		log.Println("API解析nexus数据失败: ", err)
	}
	return nexusData
}
