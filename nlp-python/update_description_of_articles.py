# This script is used to update the description field of articles because MySQL CLI in Docker didn't allow typing or pasting accented letters.
import mysql.connector
from dotenv import dotenv_values
from mysql.connector import Error

config = dotenv_values(".env")

def create_connection(host_name, username, password, db_name): 
    try:
        connection = mysql.connector.connect(
            host=host_name,
            user=username,
            passwd=password,
            database=db_name
            )
        print("Connection established.")
        return connection
    except Error as err:
        print("Error", err)
        return None

def execute_query(connection, query):
    print(query)
    cursor = connection.cursor()
    try:
        ret = cursor.execute(query)
        connection.commit()
        print("Query successful")
        return ret
    except Error as err:
        print(f"Error: '{err}'")


connection = create_connection(config['DBHOST'], config['DBUSERNAME'], config['DBPASSWORD'], config['DBNAME'])

update_desc_5 = """UPDATE article SET description="Familjen ger livet dess värde. Och inom familjen värderar vi barnen långt högre än en eventuell partner." WHERE id=5;"""
execute_query(connection, update_desc_5)

update_desc_6 = """UPDATE article SET description="När han nu väl tar plats, den traditionelle mannen, väcker han både kvinnors åtrå och mäns önskan att vara som han." WHERE id=6;"""
execute_query(connection, update_desc_6)