#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "kmeans.h"

/*
* calculate distance between two pixel points
*/
int getDistacne(Pixel *p1, Pixel *p2){
    return  ((p1->r) - (p2->r)) * ((p1->r) - (p2->r)) + 
            ((p1->g) - (p2->g)) * ((p1->g) - (p2->g)) + 
            ((p1->b) - (p2->b)) * ((p1->b) - (p2->b));
}

void generateInitCenter(Pixel* pCenter, int k){
    int i = 0;
    Pixel* max = pCenter + k;
    for (; pCenter < max; pCenter++){
        // give it random color
        pCenter->r = rand() % 255;
        pCenter->g = rand() % 255;
        pCenter->b = rand() % 255;
        // it represent the center of the ith cluster
        pCenter->clusterId = i;
        i++;
    }
}

/*
* all pixes clustering to close cluster
*/
void doClustering(Pixel* pPixes, int pixesLen, Pixel* pCenter, int centerLen){
    Pixel* pixMax = pPixes + pixesLen;
    Pixel* centerMax = pCenter + centerLen;
    for(; pPixes < pixMax; pPixes++){
        int minDistance = 999999;
        int minDistanceClusterId;
        Pixel* pc;
        // find min distance 
        for(pc = pCenter; pc < centerMax; pc++){
            int distance = getDistacne(pPixes, pc);
            if (distance < minDistance){
                minDistance = distance;
                // mark cluster id
                minDistanceClusterId = pc->clusterId;
            }
        }
        // put this pix to cluster with min distance by setting clusterId 
        pPixes->clusterId = minDistanceClusterId;
    }
}

/*
* calculate center for each cluster
*/
void recalculateCenter(Pixel* pixes, int pixesLen, Pixel* newCenter, int centerLen){

    // count of piexes in every cluster
    int* count = (int*) malloc(sizeof(int) * centerLen);
    int index;
    for (index = 0; index < centerLen; index++){
        count[index] = 0;
    }

    Pixel* maxPixes = pixes + pixesLen;
    Pixel* maxCenter = newCenter + centerLen;
    Pixel* i;
    Pixel* pCluster;
    int j = 0;

    // clear newCenter
    for(pCluster = newCenter; pCluster < maxCenter; pCluster++, j++){
        pCluster->r = 0;
        pCluster->g = 0;
        pCluster->b = 0;
        pCluster->clusterId = j;
    }

    // get sum 
    for (i = pixes; i < maxPixes; i++){
        pCluster = i->clusterId + newCenter;
        pCluster->r += i->r;
        pCluster->g += i->g;
        pCluster->b += i->b;
        count[i->clusterId]++;
    }
    
    // traverse to average 
    j = 0;
    for (i = newCenter; i < maxCenter; i++, j++){
        if (count[j] != 0){
            i->r /= count[j];
            i->g /= count[j];
            i->b /= count[j];
        }
    }
    // free memory
    free(count);
}

int isCenterStable(Pixel* oldOne, Pixel* newOne, int len){
    Pixel* maxOld = oldOne + len;
    // Pixel* maxNew = newOne + len;
    for(; oldOne < maxOld; oldOne++, newOne++){
        if (oldOne->r != newOne->r || oldOne->g != newOne->g || oldOne->b != newOne->b){
            return 0;
        }
    }
    return 1;
}

void copyCenter(Pixel* src, Pixel* dst, int len){
    Pixel* maxSrc = src + len;
    Pixel* maxDst = dst + len;
    for(; src < maxSrc; src++, dst++){
        if(src->clusterId != dst->clusterId){
            printf("error! cluster not match");
        }
        dst->r = src->r;
        dst->g = src->g;
        dst->b = src->b;
    }
}

void fillDstWithClusterCenter(Pixel* src, Pixel* dst, int pixesLen, Pixel* center){
    Pixel* max = src + pixesLen;
    for (; src < max; src++, dst++){
        dst->clusterId = src->clusterId;
        Pixel* currentCenter = center + src->clusterId;
        dst->r = currentCenter->r;
        dst->g = currentCenter->g;
        dst->b = currentCenter->b;
    }
}

/* only deal with pixes, don't mind position */
int kmeans(Pixel* src, Pixel* dst, int pixesLen, int k){

    /**************** init *****************/
    // random seed
    srand((unsigned)time(NULL));
    // prepare center array
    Pixel* center = (Pixel*) malloc(sizeof(Pixel) * k);
    generateInitCenter(center, k);

    /********** loop until converge **********/
    while(1){
        // do clustering  
        doClustering(src, pixesLen, center, k);
        // recalculate center for each cluster
        Pixel* newCenter = (Pixel*) malloc(sizeof(Pixel) * k);
        recalculateCenter(src, pixesLen, newCenter, k);
        // if the new center is close enough to the old one, break loop
        int stable = isCenterStable(center, newCenter, k);
        if (stable) {
            free(newCenter);
            break;
        } else {
            // copy new center to old one, go on next loop
            copyCenter(newCenter, center, k);
            free(newCenter);
        }
    }
    
    /********************* end processing **********************/

    fillDstWithClusterCenter(src, dst, pixesLen, center);
    // free memory
    free(center);

    // return status
    return 0;
}
