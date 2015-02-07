.PHONY: all run view edit

all: run

run:
	ansible-playbook -i inventory.ini playbook.yml --vault-password-file=$(VAULT_PASS_FILE) --limit=$(LIMIT) $(OPTS)

view:
	ansible-vault view --vault-password-file=$(VAULT_PASS_FILE) vault.yml

edit:
	ansible-vault edit --vault-password-file=$(VAULT_PASS_FILE) vault.yml