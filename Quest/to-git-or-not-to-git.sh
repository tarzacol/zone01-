returnValue=`curl -s https://learn.zone01.sn/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"mlastmock103\"}}){id}}"}'`
echo $returnValue | cut -d ':' -f 4 | cut -d '}' -f 1
