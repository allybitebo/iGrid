sudo mkdir /var/lib/postgres/data
chown postgres /var/lib/postgres/data
sudo -i -u postgres
initdb  -D '/var/lib/postgres/data'


sudo -iu postgres
initdb -D /var/lib/postgres/data
pg_ctl -D /var/lib/postgres/ -l logfile start