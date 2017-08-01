import MySQLdb
import sys
class Database:

    host = 'localhost'
    user = 'root'
    password = 'P$m7d2'
    db = 'classification'

    def __init__(self):
        self.connection = MySQLdb.connect(self.host, self.user, self.password, self.db)
        self.cursor = self.connection.cursor()
    def commit(self):
        self.connection.commit()
    def insert(self, query):
        try:
            self.cursor.execute(query)
            # print "Inserted!"
            return self.connection.insert_id()
        except:
            e = sys.exc_info()[0]
            self.connection.rollback()
            print "Failed!", query, e
            return -1


    def query(self, query):
        cursor = self.connection.cursor( MySQLdb.cursors.DictCursor )
        cursor.execute(query)

        return cursor.fetchall()

    def __del__(self):
        self.connection.close()
