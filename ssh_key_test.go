package indigo

import (
	"testing"
)

const sshKeyString1 = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC0DQty7U5izXgnzhIcgegby4EuV/BsAdb8BJZCCTFxBv5JttTV8+hd9v6XVXt+HKs2LEmRv1Bj2hw5VKV8JKVO2HBqFFRVqw4oTWPJhifboXO+WfrOy49/19nkBRVoTmK+vcRu+MaSd40vC2x8CYF0IizhOGNkJ5keKpCbllzO+nbWb7wIpr9lOevXsnAQ7fg2tihhAr3Y+CLnAJrnxHgYj9DNzB2GVbWKXeHhaPMmXIl5D6kKjdVCR7f47OXbNMp+cxUsCaT7P4dCWtyTwg2K3KFHH/Kr5oqRxJQa+SikhP0CylYTpX0fWOjLN+TjNwnvY+tAW5LXZ/h2HCZoiVkY81nda8raElV/rjBSEbpmpB0D5I7Ddaei3+4QA6BUucIxTlaKV06M+bCGroAwjfPjYt+XADm/ZHVIU7mHc0AIP2YJDB1AyRT8VXYag/xjDsbVYY/qOeYv6EHSie+h4glUdj9LjRzNZPrjIxT3CIcivle4B6QbX/CiJVy+y+aEAm0= user1@example.com"
const sshKeyString2 = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC0DQty7U5izXgnzhIcgegby4EuV/BsAdb8BJZCCTFxBv5JttTV8+hd9v6XVXt+HKs2LEmRv1Bj2hw5VKV8JKVO2HBqFFRVqw4oTWPJhifboXO+WfrOy49/19nkBRVoTmK+vcRu+MaSd40vC2x8CYF0IizhOGNkJ5keKpCbllzO+nbWb7wIpr9lOevXsnAQ7fg2tihhAr3Y+CLnAJrnxHgYj9DNzB2GVbWKXeHhaPMmXIl5D6kKjdVCR7f47OXbNMp+cxUsCaT7P4dCWtyTwg2K3KFHH/Kr5oqRxJQa+SikhP0CylYTpX0fWOjLN+TjNwnvY+tAW5LXZ/h2HCZoiVkY81nda8raElV/rjBSEbpmpB0D5I7Ddaei3+4QA6BUucIxTlaKV06M+bCGroAwjfPjYt+XADm/ZHVIU7mHc0AIP2YJDB1AyRT8VXYag/xjDsbVYY/qOeYv6EHSie+h4glUdj9LjRzNZPrjIxT3CIcivle4B6QbX/CiJVy+y+aEAm0= user2@example.com"
const sshKeyString3 = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC0DQty7U5izXgnzhIcgegby4EuV/BsAdb8BJZCCTFxBv5JttTV8+hd9v6XVXt+HKs2LEmRv1Bj2hw5VKV8JKVO2HBqFFRVqw4oTWPJhifboXO+WfrOy49/19nkBRVoTmK+vcRu+MaSd40vC2x8CYF0IizhOGNkJ5keKpCbllzO+nbWb7wIpr9lOevXsnAQ7fg2tihhAr3Y+CLnAJrnxHgYj9DNzB2GVbWKXeHhaPMmXIl5D6kKjdVCR7f47OXbNMp+cxUsCaT7P4dCWtyTwg2K3KFHH/Kr5oqRxJQa+SikhP0CylYTpX0fWOjLN+TjNwnvY+tAW5LXZ/h2HCZoiVkY81nda8raElV/rjBSEbpmpB0D5I7Ddaei3+4QA6BUucIxTlaKV06M+bCGroAwjfPjYt+XADm/ZHVIU7mHc0AIP2YJDB1AyRT8VXYag/xjDsbVYY/qOeYv6EHSie+h4glUdj9LjRzNZPrjIxT3CIcivle4B6QbX/CiJVy+y+aEAm0= user3@example.com"

var sshKeyForTestSSHKey *SSHKey

func TestCreateSSHKey(t *testing.T) {
	var err error
	sshKeyForTestSSHKey, err = client.CreateSSHKey(
		"sshKeyForTestSSHKey",
		sshKeyString2,
	)
	if sshKeyForTestSSHKey == nil {
		t.Fatalf("CreateSSHKey() = %v, want %v (%v)", sshKeyForTestSSHKey, "'not nil'", err)
	}
}

func TestGetSSHKeyList(t *testing.T) {
	keys, err := client.GetSSHKeyList()
	if err != nil || len(keys) != 2 {
		t.Fatalf("GetSSHKeyList() = %v, want %v (%v)", err, "'not nil'", keys)
	}
}

func TestGetActiveSSHKeyList(t *testing.T) {
	keys, err := client.GetActiveSSHKeyList()
	if err != nil || len(keys) != 2 {
		t.Fatalf("GetActiveSSHKeyList() = %v, want %v (%v)", err, "'not nil'", keys)
	}
}

func TestRetrieveSSHKey(t *testing.T) {
	key, err := client.RetrieveSSHKey(sshKeyForTestSSHKey.ID)
	if key == nil {
		t.Fatalf("RetrieveSSHKey() = %v, want %v (%v)", key, "'not nil'", err)
	}
}

func TestUpdateSSHKey(t *testing.T) {
	err := client.UpdateSSHKey(sshKeyForTestSSHKey.ID, "updatedKey", sshKeyString3, "DEACTIVE")
	if err != nil {
		t.Fatalf("UpdateSSHKey() = %v, want %v", err, "'nil'")
	}
	key, err := client.RetrieveSSHKey(sshKeyForTestSSHKey.ID)
	if err != nil {
		t.Fatalf("RetrieveSSHKey() = %v, want %v", err, "nil")
	}
	if key.Name != "updatedKey" || key.Key != sshKeyString3 || key.Status != "DEACTIVE" {
		t.Fatalf("RetrieveSSHKey() = %v, want {%s, %s, %s}", key, "updatedKey", sshKeyString3, "DEACTIVE")
	}
}

func TestDeleteSSHKey(t *testing.T) {
	err := client.DeleteSSHKey(sshKeyForTestSSHKey.ID)
	if err != nil {
		t.Fatalf("DeleteSSHKey() = %v, want %v", err, "'nil'")
	}
}
