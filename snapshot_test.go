package indigo

import (
	"fmt"
	"testing"
)

var snapshotsForTest []*Snapshot

func TestCreateSnapshot(t *testing.T) {
	err := client.CreateSnapshot("snapshot1", instanceForTest.ID)
	if err != nil {
		t.Fatalf("CreateSnapshot() = %v, want %v", err, "nil")
	}
}

func TestCreateSnapshotSync(t *testing.T) {
	snapshot, err := client.CreateSnapshotSync("snapshot1", instanceForTest.ID)
	if err != nil {
		t.Fatalf("CreateSnapshot() = %v, want %v", err, "nil")
	}
	fmt.Printf("%v\n", snapshot)
}

func TestGetSnapshotList(t *testing.T) {
	var err error
	snapshotsForTest, err = client.GetSnapshotList(instanceForTest.ID)
	if err != nil {
		t.Fatalf("GetSnapshotList() = %v, want %v", err, "nil")
	}
	if len(snapshotsForTest) == 0 {
		t.Fatalf("GetSnapshotList() = %v, want %v", snapshotsForTest, "'not empty'")
	}
}

func TestRecreateSnapshot(t *testing.T) {
	err := client.RecreateSnapshot(instanceForTest.ID, snapshotsForTest[0].ID)
	if err != nil {
		t.Fatalf("RecreateSnapshot() = %v, want %v", err, "nil")
	}
}

func TestRestoreSnapshot(t *testing.T) {
	err := client.RestoreSnapshot(instanceForTest.ID, snapshotsForTest[0].ID)
	if err != nil {
		t.Fatalf("RestoreSnapshot() = %v, want %v", err, "nil")
	}
}

func TestDeleteSnapshot(t *testing.T) {
	err := client.DeleteSnapshot(snapshotsForTest[0].ID)
	if err != nil {
		t.Fatalf("DeleteSnapshot() = %v, want %v", err, "nil")
	}
}
