package indigo

import (
	"testing"
	"time"
)

const sshKeyString1 = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC0DQty7U5izXgnzhIcgegby4EuV/BsAdb8BJZCCTFxBv5JttTV8+hd9v6XVXt+HKs2LEmRv1Bj2hw5VKV8JKVO2HBqFFRVqw4oTWPJhifboXO+WfrOy49/19nkBRVoTmK+vcRu+MaSd40vC2x8CYF0IizhOGNkJ5keKpCbllzO+nbWb7wIpr9lOevXsnAQ7fg2tihhAr3Y+CLnAJrnxHgYj9DNzB2GVbWKXeHhaPMmXIl5D6kKjdVCR7f47OXbNMp+cxUsCaT7P4dCWtyTwg2K3KFHH/Kr5oqRxJQa+SikhP0CylYTpX0fWOjLN+TjNwnvY+tAW5LXZ/h2HCZoiVkY81nda8raElV/rjBSEbpmpB0D5I7Ddaei3+4QA6BUucIxTlaKV06M+bCGroAwjfPjYt+XADm/ZHVIU7mHc0AIP2YJDB1AyRT8VXYag/xjDsbVYY/qOeYv6EHSie+h4glUdj9LjRzNZPrjIxT3CIcivle4B6QbX/CiJVy+y+aEAm0= user1@example.com"
const sshKeyString2 = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC0DQty7U5izXgnzhIcgegby4EuV/BsAdb8BJZCCTFxBv5JttTV8+hd9v6XVXt+HKs2LEmRv1Bj2hw5VKV8JKVO2HBqFFRVqw4oTWPJhifboXO+WfrOy49/19nkBRVoTmK+vcRu+MaSd40vC2x8CYF0IizhOGNkJ5keKpCbllzO+nbWb7wIpr9lOevXsnAQ7fg2tihhAr3Y+CLnAJrnxHgYj9DNzB2GVbWKXeHhaPMmXIl5D6kKjdVCR7f47OXbNMp+cxUsCaT7P4dCWtyTwg2K3KFHH/Kr5oqRxJQa+SikhP0CylYTpX0fWOjLN+TjNwnvY+tAW5LXZ/h2HCZoiVkY81nda8raElV/rjBSEbpmpB0D5I7Ddaei3+4QA6BUucIxTlaKV06M+bCGroAwjfPjYt+XADm/ZHVIU7mHc0AIP2YJDB1AyRT8VXYag/xjDsbVYY/qOeYv6EHSie+h4glUdj9LjRzNZPrjIxT3CIcivle4B6QbX/CiJVy+y+aEAm0= user2@example.com"

func TestCreateSSHKey(t *testing.T) {
	/*
		time.Sleep(time.Second * 8)
		key, err := client.CreateSSHKey(
			"testkey",
			sshKeyString1,
		)
		if key == nil {
			t.Fatalf("CreateSSHKey() = %v, want %v (%v)", key, "'not nil'", err)
		}
	*/
}

func TestGetSSHKeyList(t *testing.T) {
	time.Sleep(time.Second * 8)
	keys, err := client.GetSSHKeyList()
	if err != nil || len(keys) != 1 {
		t.Fatalf("GetSSHKeyList() = %v, want %v (%v)", err, "'not nil'", keys)
	}
}

func TestGetActiveSSHKeyList(t *testing.T) {
	time.Sleep(time.Second * 8)
	keys, err := client.GetActiveSSHKeyList()
	if err != nil || len(keys) != 1 {
		t.Fatalf("GetActiveSSHKeyList() = %v, want %v (%v)", err, "'not nil'", keys)
	}
}

func TestRetrieveSSHKey(t *testing.T) {
	time.Sleep(time.Second * 8)
	retKey, err := client.RetrieveSSHKey(key.ID)
	if retKey == nil {
		t.Fatalf("RetrieveSSHKey() = %v, want %v (%v)", retKey, "'not nil'", err)
	}
}

func TestUpdateSSHKey(t *testing.T) {
	time.Sleep(time.Second * 8)
	err := client.UpdateSSHKey(key.ID, "updatedtestkey", sshKeyString2, "DEACTIVE")
	if err != nil {
		t.Fatalf("UpdateSSHKey() = %v, want %v", err, "'nil'")
	}
	time.Sleep(time.Second * 8)
	retKey, err := client.RetrieveSSHKey(key.ID)
	if retKey.Name != "updatedtestkey" || retKey.Key != sshKeyString2 || retKey.Status != "DEACTIVE" {
		t.Fatalf("RetrieveSSHKey() = %v, want {%s, %s, %s}", retKey, "updatedtestkey", sshKeyString2, "DEACTIVE")
	}
}

func TestDestroySSHKey(t *testing.T) {
	time.Sleep(time.Second * 8)
	err := client.DestroySSHKey(key.ID)
	if err != nil {
		t.Fatalf("DestroySSHKey() = %v, want %v", err, "'nil'")
	}
}
