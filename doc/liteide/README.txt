LITEIDE (27.2.1)

1)
�������� � ���������������� ���� ������ ��� �� ��������� �������
(�������� ����� ���������/share/liteide/litebuild/gosrc.xml)

		<action id="Install1" menu="Build" img="gray/install.png" key="Ctrl+1" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP1)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install2" menu="Build" img="gray/install.png" key="Ctrl+2" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP2)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install3" menu="Build" img="gray/install.png" key="Ctrl+3" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP3)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install4" menu="Build" img="gray/install.png" key="Ctrl+4" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP4)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install5" menu="Build" img="gray/install.png" key="Ctrl+5" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP5)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install6" menu="Build" img="gray/install.png" key="Ctrl+6" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP6)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install7" menu="Build" img="gray/install.png" key="Ctrl+7" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP7)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install8" menu="Build" img="gray/install.png" key="Ctrl+8" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP8)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>
		<action id="Install9" menu="Build" img="gray/install.png" key="Ctrl+9" cmd="$(GO)" args="install -a $(GOPATH)/src/$(PATHAPP9)" save="all" output="true" codec="utf-8" regex="$(ERRREGEX)" navigate="true"/>

��� ����������� ����� ���������� � ��������� args
���������� ��� �������� ���������� ���������� ����������� ����������
"Ctrl+9" ��� ������� ������� �� ������ ����� ���������� ������������ ����������
		
2) 
�������� � ���������������� ���� ���������� ��������� ��������� �������
(�������� ����� ���������/share/liteide/liteenv/win64.env)

PATHAPP1=application.go
PATHAPP2=application.go
PATHAPP3=application.go
PATHAPP4=application.go
PATHAPP5=application.go
PATHAPP6=application.go
PATHAPP7=application.go
PATHAPP8=application.go
PATHAPP9=application.go

������ ��� ������ ����� ��� ��������� ������� ������ (����������) ����� �������������� ���������� ������� �� ������� �������� �� ������ ��������� �����.
"application.go" ��� ���������� ��� �������� ����� � ������ ����� � ���������� ���������� ������� � ����� ����������� (src)
������� ���� ����� ���������� ������� �� ����������.
��� ������ �� ���������� ����� ������������ 9. ����� ��� ����������.
�� ���� ����� ������, ������ ��������� ���������������� ������� ������� ������.
������� ����� ��� ���� ���������������� ����� 1 � 2 �������

3)
�� �������� ���������������� ���������� GOBIN � ��������� ���� �� ����� ���� ����� ���������� ���������. �� ��������� ��� ���������������.
� ����� ���������� ���������� GOPATH � ���������� ��������� IDE

4) 
Keys.kms
��� ��� �������� ������� ������ ��� �������� ������
���� ��� ���������...

5)
��� ��������� ��������� �� �����
