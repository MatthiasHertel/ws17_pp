# Entwicklung einer Security Suite


## Motivation

Als Administratoren von HPC Clustern haben wir das Problem, dass wir auf bekannt gewordene Sicherheitslücken in der eingesetzten Software reagieren müssen. Dies kann aber nicht immer durch ein zeitnahes Update geschehen, denn die bei uns berechneten HPC Jobs können so gross sein, dass sie auf mehreren Servern mehrere Wochen rechnen. Diese Server kann man in dieser Zeit nicht updaten.

Im Security Board werden Sicherheitsthemen regelmäßig angesprochen, aber der manuelle Aufwand ist mittlerweile zu groß um relevante von irrelevanten Sicherheitslücken zu unterscheiden und verfolgen zu können, welche Security-Bugs auf welchen Servern behoben wurden bzw. noch zu beheben sind.

Das Tool soll wichtige Unterstützung für diese Prozesse liefern und die Informationen so auswerten, dass sie Umsetzbar werden.


## Beschreibung

Um den Admins einen besseren Überblick zu verschaffen und auch zur Kommunikation mit den Kunden, benötigen wir einen Abgleich, welche Software Versionen auf welchem Server, Docker Container, Virtuellen Maschine oder Firewall installiert ist und damit möglicherweise relevant für Sicherheitslücken sind.

Die Anforderungen können grob in zwei Fälle unterteilt werden:

Software wurde mittels Paketverwaltung installiert
Paketverwaltungen wie apt, yum, pip, ruby gems ermöglichen eine einfache Installation von Software. Die aktuellen Versionsstände der installierten Software kann man auf einem einzelnen System schnell abfragen und sollte zentral gespeichert werden.

Software wurde ohne Paketverwaltung installiert
In diesem Fall sollte es eine Möglichkeit geben, den Agenten so zu konfigurieren, dass Abfragen leicht möglich sein. Beispielsweise koennte das ueber ein Kommando wie `<software> --version` passieren, oder der aktuellen git commit hash indentifiziert die entsprechende Softwareversion. Dies soll möglichst einfach im Terminal eingepflegt werden können. Als Beispiel soll z.B. nach einem Git Checkout der aktuelle Commit Hash (plus Hostname, Installationsverzeichnis und Zeitstempel) erfasst werden.

Diese Informationen sollen an einer zentralen Stelle/Repository im Cluster hinterlegt werden. Die Zusammenfassung der Installation über Cluster-Grenzen hinweg soll möglich sein.

Es soll möglich sein, Abfragen über den Datenbestand zu treffen und dabei auch historische Daten zu betrachten. Dazu müssen historische Daten, insbesondere Änderungen gespeichert werden.

Abfragen könnten sein:

Welche Version einer bestimmten Software sind wo im Cluster installiert?
Welche Versionen sind bezüglich einer Sicherheitslücke relevant und wo sind diese installiert?
In welchem Zeitraum war auf welchem Server welche Software-Version installiert?


## Umsetzung

Die Software soll automatisch relevante Security-Feeds importieren können. Beispielsweise die Veroeffentlichungen unter https://usn.ubuntu.com/usn/

CVE-Nummern und andere Identifikatoren, Software Pakete, relevante Versionen, Beschreibung und weitere Meta-Daten sollen erfasst werden. https://people.canonical.com/~ubuntu-security/cve/2017/CVE-2017-13721.html

Regelmäßig (mindestens täglich) sollen die installierten Versionen aller Pakete zentral im Cluster gespeichert werden. Dazu muss ein Agent Zugriff auf die Paketverwaltungen (z.B. apt, yum, zypper, pip, Ruby Gems, Docker) haben. Es ist sicherlich sinnvoll, wenn der Agent selbst auch auf der Maschine (Server, VM, Docker Container) installiert ist, um auch manuell Versionsstände speichern zu können ohne die Maschine verlassen zu müssen.
d
Regelmäßig (täglich oder wöchentlich) sollen an zentraler Stelle im Cluster Reports erstellt und mittels E-Mail verschickt werden. Diese Reports sollen auch manuell erstellt und verschickt werden können.

Es soll möglich sein, eine Liste mit Servern zu erhalten, die mittels Ansible bezüglich einer bestimmten Lücke gepatcht werden sollen. Dazu wäre eine Integration als Ansible Inventory-Plugin hilfreich.

http://docs.ansible.com/ansible/latest/intro_dynamic_inventory.html
http://docs.ansible.com/ansible/latest/dev_guide/developing_inventory.html

## Vorgehensweise

In iterativen Entwicklungszyklen von ca 2 Wochen sollen sowohl Command Line Interface (CLI) als auch Rest API in Go oder Python entwickelt werden. Während der Entwicklungszeit soll das Tool regelmäßig vom “Kunden” (den Administratoren der HPC Cluster) getestet werden und Verbesserungsvorschläge umgesetzt werden.

Da ein geeignetes Datenmodell und eventuelle Probleme vermutlich erst mit dem Verlauf des Projektes erkennbar werden, kann in den ersten Phasen der Umsetzung darauf verzichtet werden, historische Daten zu migrieren. Diese liegen als regelmäßige Reports eh vor. Wichtiger ist ein möglichst umfassender Blick und ein Abgleich mit möglichst vielen Informationsquellen, ohne dass die Flut der Informationen zu gross wird.
