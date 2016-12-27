package search

import (
	"fmt"

	swanclient "github.com/Dataman-Cloud/swan-search/src/util/go-swan"
	log "github.com/Sirupsen/logrus"
)

const (
	DOCUMENT_APP       = "app"
	DOCUMENT_CONTAINER = "container"
	DOCUMENT_TASK      = "task"
)

type SwanIndexer struct {
	Indexer

	SwanClients []swanclient.Swan
}

func NewSwanIndex(SwanClients []swanclient.Swan) *SwanIndexer {
	return &SwanIndexer{SwanClients: SwanClients}
}

func (indexer *SwanIndexer) Index(store *DocumentStorage) {
	for _, swanClient := range indexer.SwanClients {
		var filter map[string][]string
		if apps, err := swanClient.Applications(filter); err == nil {
			fmt.Printf("applications:%+v", apps)
			for _, app := range apps {
				store.Set(app.ID+app.Name, Document{
					ID:   app.ID,
					Name: app.Name,
					Type: DOCUMENT_APP,
					Param: map[string]string{
						"AppId": app.ID,
					},
				})
			}
		} else {
			log.Warnf("get applications error:", err)
		}
	}
	//if nodes, err := indexer.SwanDockerClient.ListNode(types.NodeListOptions{}); err == nil {
	//	for _, node := range nodes {
	//		store.Set(node.ID+node.Description.Hostname, Document{
	//			ID:   node.ID,
	//			Name: node.Description.Hostname,
	//			Type: DOCUMENT_NODE,
	//			Param: map[string]string{
	//				"NodeId": node.ID,
	//			},
	//		})

	//		backContext := context.WithValue(context.Background(), "node_id", node.ID)
	//		if networks, err := indexer.
	//			SwanDockerClient.
	//			ListNodeNetworks(backContext, docker.NetworkFilterOpts{}); err == nil {
	//			for _, network := range networks {
	//				store.Set(network.ID+network.Name+node.ID, Document{
	//					Name: network.Name,
	//					ID:   network.ID,
	//					Type: DOCUMENT_NETWORK,
	//					Param: map[string]string{
	//						"NodeId":    node.ID,
	//						"NetworkID": network.ID,
	//					},
	//				})

	//			}
	//		} else {
	//			log.Warnf("get network error: %v", err)
	//		}

	//		if volumes, err := indexer.
	//			SwanDockerClient.
	//			ListVolumes(backContext, docker.ListVolumesOptions{}); err == nil {
	//			for _, volume := range volumes {
	//				store.Set(volume.Name, Document{
	//					Name: volume.Name,
	//					Type: DOCUMENT_VOLUME,
	//					Param: map[string]string{
	//						"NodeId":     node.ID,
	//						"VolumeName": volume.Name,
	//					},
	//				})
	//			}
	//		} else {
	//			log.Warnf("get volume error: %v", err)
	//		}
	//	}
	//} else {
	//	log.Warnf("get node list error: %v", err)
	//}

	//if stacks, err := indexer.SwanDockerClient.ListStack(); err == nil {
	//	for _, stack := range stacks {
	//		//bundle, _ := indexer.Swandockerclient.InspectStack(stack.Namespace)
	//		//groupId, _ := indexer.SwanDockerClient.GetStackGroup(bundle)
	//		groupId := uint64(1)

	//		store.Set(stack.Namespace, Document{
	//			ID:      stack.Namespace,
	//			Type:    DOCUMENT_STACK,
	//			GroupId: groupId,
	//			Param: map[string]string{
	//				"NameSpace": stack.Namespace,
	//			},
	//		})

	//		if services, err := indexer.
	//			SwanDockerClient.
	//			ListStackService(stack.Namespace, types.ServiceListOptions{}); err == nil {
	//			for _, service := range services {
	//				store.Set(service.ID+stack.Namespace,
	//					Document{
	//						ID:      service.ID,
	//						Name:    stack.Namespace,
	//						Type:    DOCUMENT_SERVICE,
	//						GroupId: groupId,
	//						Param: map[string]string{
	//							"NameSpace": stack.Namespace,
	//							"ServiceId": service.ID,
	//						},
	//					})

	//				if tasks, err := indexer.
	//					SwanDockerClient.
	//					ListTasks(types.TaskListOptions{}); err == nil {
	//					for _, task := range tasks {
	//						store.Set(task.ID,
	//							Document{
	//								ID:      task.ID,
	//								Type:    DOCUMENT_TASK,
	//								GroupId: groupId,
	//								Param: map[string]string{
	//									"NodeId":      task.NodeID,
	//									"ContainerId": task.Status.ContainerStatus.ContainerID,
	//								},
	//							})
	//						store.Set(task.Status.ContainerStatus.ContainerID,
	//							Document{
	//								ID:      task.Status.ContainerStatus.ContainerID,
	//								Type:    DOCUMENT_TASK,
	//								GroupId: groupId,
	//								Param: map[string]string{
	//									"NodeId":      task.NodeID,
	//									"ContainerId": task.Status.ContainerStatus.ContainerID,
	//								},
	//							})
	//					}
	//				} else {
	//					log.Warnf("get task list error: %v", err)
	//				}
	//			}
	//		} else {
	//			log.Warnf("get service error: %v", err)
	//		}

	//	}
	//} else {
	//	log.Warnf("get stack list error: %v", err)
	//}
}
