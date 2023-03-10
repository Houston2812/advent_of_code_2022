package tree

import (
	"fmt"
	"strings"
)

type EntityDescriptor struct {
	Name  string
	Type_ string
	Size  int
}

func PrintEntity(entity EntityDescriptor) {
	// if entity.Type_ == "dir" {
		
		fmt.Printf("[-] Name: %s; Type: %s; Size: %d;\n", entity.Name, entity.Type_, entity.Size)
	// }	 
	

}

func compareEntities(entity1 EntityDescriptor, entity2 EntityDescriptor) bool {
	if entity1.Name == entity2.Name && entity1.Size == entity2.Size &&  entity1.Type_ == entity2.Type_ {
		return true
	} else {
		return false
	}
}

type FileSystemEntity struct {
	entity EntityDescriptor
	is_root bool
	sub_entities []*FileSystemEntity
}

type FileSystem struct {
	Root *FileSystemEntity
}

func (fileSystem *FileSystem) Insert(entity EntityDescriptor, parent ...EntityDescriptor) {
	fileSystem.InsertRec(fileSystem.Root, entity, parent...)
}


func (fileSystem *FileSystem) InsertRec(fileSystemEntity *FileSystemEntity, entity EntityDescriptor, parent ...EntityDescriptor) *FileSystemEntity {
	
	if fileSystem.Root == nil && parent == nil {
		fileSystem.Root = &FileSystemEntity{entity, true, []*FileSystemEntity {}	}
		return fileSystem.Root
	}
	
	if compareEntities(fileSystemEntity.entity, parent[0]) {
		fileSystemEntity.sub_entities = append(fileSystemEntity.sub_entities, &FileSystemEntity{entity, false, make([]*FileSystemEntity,0)})
		return fileSystemEntity
	}

	if len(fileSystemEntity.sub_entities) == 0 {
		return fileSystemEntity
	}

	
	for _, ent := range fileSystemEntity.sub_entities {
		ent = fileSystem.InsertRec(ent, entity, parent[0])
		_ = ent
		// PrintEntity(ent.entity)
	}

	return fileSystemEntity
}

func (fileSystem *FileSystem) PrintInorder(fileSystemEntity *FileSystemEntity, level int) {
	if fileSystemEntity == nil {
		return
	} else {

		if fileSystemEntity.is_root {
			indent := strings.Repeat("\t", level)
			fmt.Printf("%s", indent)
			
			PrintEntity(fileSystemEntity.entity)
			
			if (fileSystemEntity.entity.Type_ == "dir"){

				level++
			}
		}

		for _, entity := range fileSystemEntity.sub_entities {
			// if entity.entity.Type_ == "dir"{
				level++
			// }

			indent := strings.Repeat("\t", level)
			fmt.Printf("%s", indent)

			PrintEntity(entity.entity)

			if entity.entity.Type_ == "dir"{

				fileSystem.PrintInorder(entity, level)
			}
			
			// if entity.entity.Type_ == "dir" {
				level--
			// }

		}
		return
	}
}

func (fileSystem *FileSystem) GetTotalSize(fileSystemEntity *FileSystemEntity ) int {
	if fileSystemEntity == nil {
		return 0
	} else {
		var size int = 0
		for _, entity := range fileSystemEntity.sub_entities {
			if entity.entity.Type_ == "file" {
				size += entity.entity.Size
			} else {
				// var size_tmp int = 0

				size  += fileSystem.GetTotalSize(entity)
				// entity.entity.Size = size_tmp
				
				// size += size_tmp
			}
		}
		return size
	}
}

func (fileSystem *FileSystem) Search(entity EntityDescriptor) bool {
    found := fileSystem.SearchRec(fileSystem.Root, entity)
    return found
}

func (fileSystem *FileSystem) SearchRec(fileSystemEntity *FileSystemEntity , entity EntityDescriptor) bool {
    if fileSystemEntity == nil {
        return false
    }

    if compareEntities(fileSystemEntity.entity, entity) {
        return true
    }

	for _, ent := range fileSystemEntity.sub_entities {
		is_found := fileSystem.SearchRec(ent, entity)
		
		if is_found {
			return true
		}
		// PrintEntity(ent.entity)
	}


    return false
}


// func main() {
// 	fileSystem := FileSystem{}

// 	fmt.Printf("Started\n")
// 	fileSystem.Insert(EntityDescriptor{"a", "dir", 24})

// 	fmt.Printf("Inserted\n")

// 	parent := make([]EntityDescriptor, 0)
// 	parent = append(parent,  EntityDescriptor{"a", "dir", 24})

// 	fileSystem.Insert(EntityDescriptor{"b", "file", 25}, parent...)
// 	fmt.Printf("Inserted\n")

// 	fileSystem.Insert(EntityDescriptor{"c", "file", 26}, parent...)
// 	fmt.Printf("Inserted\n")

// 	fileSystem.Insert(EntityDescriptor{"d", "file", 27}, parent...)
// 	fmt.Printf("Inserted\n")

// 	fileSystem.Insert(EntityDescriptor{"e", "file", 28}, parent...)
// 	fmt.Printf("Inserted\n")


// 	parent[0] = EntityDescriptor{"c", "file", 26}

// 	fileSystem.Insert(EntityDescriptor{"h", "file", 34}, parent...)
// 	fmt.Printf("Printing\n")

// 	result := fileSystem.Search(EntityDescriptor{"h", "file", 34})

// 	fmt.Printf("The entity is found: %v\n", result)

// 	fileSystem.PrintInorder(fileSystem.Root, 0)
// }
